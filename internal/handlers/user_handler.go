package handlers

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mhcodev/fake_store_api/internal/models"
	"github.com/mhcodev/fake_store_api/internal/services"
	"github.com/mhcodev/fake_store_api/internal/util"
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
	limit, err := strconv.Atoi(c.Query("limit", "0"))

	if err != nil {
		limit = 15
	}

	offset, err := strconv.Atoi(c.Query("offset", "0"))

	if err != nil {
		offset = 0
	}

	params.Limit = limit
	params.Offset = offset

	users, err := h.UserService.GetUsersByParams(c.Context(), params)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Return user as JSON
	return c.JSON(users)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	// Call the service to get user by id
	userID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	users, err := h.UserService.GetUserByID(c.Context(), userID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Return user as JSON
	return c.JSON(users)
}

type UserEmailIsAvailableRequest struct {
	Email string `json:"email"`
}

func (h *UserHandler) UserEmailIsAvailable(c *fiber.Ctx) error {
	var request UserEmailIsAvailableRequest
	messages := make([]string, 0)

	// Parse the request body into the struct
	if err := c.BodyParser(&request); err != nil {
		messages = append(messages, "Error processing request")
		return util.ErrorReponse(c, fiber.StatusInternalServerError, nil, messages)
	}

	if request.Email == "" {
		messages = append(messages, "email is empty, ensure you send an email")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	email := strings.ToLower(request.Email)

	if email == "none" {
		messages = append(messages, "email is empty, ensure you send an email")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	response, err := h.UserService.UserEmailIsAvailable(c.Context(), email)

	if err != nil {
		messages = append(messages, err.Error())
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	// Return user as JSON
	return util.SuccessReponse(c, response)
}

type CreateUserRequest struct {
	User models.User `json:"user"`
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var request CreateUserRequest
	messages := make([]string, 0)

	// Parse the request body into the struct
	if err := c.BodyParser(&request); err != nil {
		messages = append(messages, "Error processing request")
		return util.ErrorReponse(c, fiber.StatusInternalServerError, nil, messages)
	}

	user := request.User

	typesAvailable := make([]int, 0)
	typesAvailable = append(typesAvailable, 1, 2, 3)

	if !util.Includes(typesAvailable, user.UserTypeID) {
		messages = append(messages, "user type id is not valid")
	}

	err := h.UserService.CreateUser(c.Context(), &user)

	if err != nil {
		messages = append(messages, err.Error())
		return util.ErrorReponse(c, fiber.StatusInternalServerError, nil, messages)
	}

	response := make(map[string]interface{})
	response["user"] = user

	// Return user as JSON
	return util.SuccessReponse(c, response)
}
