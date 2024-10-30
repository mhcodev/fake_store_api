package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mhcodev/fake_store_api/internal/models"
	"github.com/mhcodev/fake_store_api/internal/services"
	"github.com/mhcodev/fake_store_api/internal/util"
	"github.com/mhcodev/fake_store_api/internal/validators"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) GetUsersByParams(c *fiber.Ctx) error {
	// Call the service to get user details
	var params models.QueryParams

	limit := c.QueryInt("limit", 15)
	offset := c.QueryInt("offset", 0)

	name := c.Query("name")
	userTypeID := c.QueryInt("type", -1)
	email := c.Query("email")
	status := c.QueryInt("status", -1)

	params.MapParams = make(map[string]interface{}, 0)
	params.MapParams["limit"] = limit
	params.MapParams["offset"] = offset
	params.MapParams["name"] = name
	params.MapParams["type"] = userTypeID
	params.MapParams["email"] = email
	params.MapParams["status"] = status

	params.Limit = limit
	params.Offset = offset

	users, err := h.UserService.GetUsersByParams(c.Context(), params)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response := make(map[string]interface{}, 0)
	response["users"] = users
	return util.SuccessReponse(c, response)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	// Call the service to get user by id
	userID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.UserService.GetUserByID(c.Context(), userID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Return user as JSON
	response := make(map[string]interface{}, 0)
	response["user"] = user
	return util.SuccessReponse(c, response)
}

type UserEmailIsAvailableRequest struct {
	Email string `json:"email"`
}

func (h *UserHandler) UserEmailIsAvailable(c *fiber.Ctx) error {
	var request UserEmailIsAvailableRequest
	var validationErrors = validators.ValidationErrors{}

	// Parse the request body into the struct
	if err := c.BodyParser(&request); err != nil {
		validationErrors.AddError("msg", "Error processing request")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	if request.Email == "" {
		validationErrors.AddError("msg", "email is empty, ensure you send an email")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	email := strings.ToLower(request.Email)

	if email == "none" {
		validationErrors.AddError("msg", "email is empty, ensure you send an email")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	response, err := h.UserService.UserEmailIsAvailable(c.Context(), email)

	if err != nil {
		validationErrors.AddError("msg", err.Error())
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	return util.SuccessReponse(c, response)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var input services.UserCreateInput
	var validationErrors = validators.ValidationErrors{}

	// Parse the request body into the struct
	if err := c.BodyParser(&input); err != nil {
		validationErrors.AddError("msg", "Error processing request")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	if validators.IsNotEmpty(input.Email) && !validators.IsValidEmail(*input.Email) {
		validationErrors.AddError("email", "email is not valid")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	validationErrors = validators.ValidateUserCreateInput(input)

	if validationErrors.HasErrors() {
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	emailIsUsed, err := h.UserService.UserEmailIsAvailable(c.Context(), *input.Email)

	if err != nil || emailIsUsed["isAvailable"] == false {
		validationErrors.AddError("email", "email is already used")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	user, err := h.UserService.CreateUser(c.Context(), input)

	if err != nil {
		validationErrors.AddError("msg", "user no created, check your body request")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	response := make(map[string]interface{})
	response["user"] = user

	// Return user as JSON
	return util.SuccessReponse(c, response)
}

type UpdateUserRequest struct {
	User models.User `json:"user"`
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	var input services.UserUpdateInput
	var validationErrors = validators.ValidationErrors{}

	// Parse the request body into the struct
	if err := c.BodyParser(&input); err != nil {
		validationErrors.AddError("msg", "error processing request")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	userID, err := c.ParamsInt("id", 0)

	if err != nil || userID == 0 {
		validationErrors.AddError("msg", "user id is not valid")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	_, err = h.UserService.GetUserByID(c.Context(), userID)

	if err != nil {
		validationErrors.AddError("msg", "user not found")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	validationErrors = validators.ValidateUserUpdateInput(input)

	userUpdated, err := h.UserService.UpdateUser(c.Context(), userID, input)

	if err != nil {
		validationErrors.AddError("msg", err.Error())
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	response := make(map[string]interface{})
	response["user"] = userUpdated

	// Return user as JSON
	return util.SuccessReponse(c, response)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id", "0"))
	var validationErrors = validators.ValidationErrors{}

	if err != nil {
		validationErrors.AddError("msg", "user id is not valid")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	_, err = h.UserService.GetUserByID(c.Context(), userID)

	if err != nil {
		validationErrors.AddError("msg", "user doesn't exist")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	err = h.UserService.DeletedUser(c.Context(), userID)

	if err != nil {
		validationErrors.AddError("msg", err.Error())
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	response := make(map[string]interface{})
	response["msg"] = fmt.Sprintf("user %d deleted", userID)

	// Return user as JSON
	return util.SuccessReponse(c, response)
}
