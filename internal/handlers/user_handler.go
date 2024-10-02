package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mhcodev/fake_store_api/internal/models"
	"github.com/mhcodev/fake_store_api/internal/services"
	"github.com/mhcodev/fake_store_api/internal/util"
	"golang.org/x/crypto/bcrypt"
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
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
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
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	user := request.User

	userTypes, err := h.UserService.GetUserTypes(c.Context())

	if err != nil {
		messages = append(messages, "user types no available")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	var typesAvailable []int

	for _, userType := range userTypes {
		typesAvailable = append(typesAvailable, userType.ID)
	}

	if len(typesAvailable) > 0 && !util.Includes(typesAvailable, user.UserTypeID) {
		messages = append(messages, "user type id is not valid")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	if strings.TrimSpace(request.User.Name) == "" {
		messages = append(messages, "name is required")
	}

	if strings.TrimSpace(request.User.Email) == "" {
		messages = append(messages, "email is required")
	} else {
		response, _ := h.UserService.UserEmailIsAvailable(c.Context(), strings.TrimSpace(request.User.Email))

		if response["isAvailable"] == false {
			messages = append(messages, "email is already used")
		}
	}

	if strings.TrimSpace(request.User.Password) == "" {
		messages = append(messages, "password is required")
	} else {
		password := strings.TrimSpace(request.User.Password)
		passwordHashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)

		if err != nil {
			messages = append(messages, "password is not valid")
			return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
		}

		user.Password = string(passwordHashed)
	}

	if strings.TrimSpace(request.User.Avatar) != "" {
		isImage, err := util.IsImageURL(strings.TrimSpace(request.User.Avatar))
		if err != nil {
			messages = append(messages, "avatar has to be a valid image")
		} else if !isImage {
			messages = append(messages, "avatar has to be a valid image")
		} else {
			user.Avatar = strings.TrimSpace(request.User.Avatar)
		}
	}

	if len(messages) > 0 {
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	err = h.UserService.CreateUser(c.Context(), &user)

	if err != nil {
		messages = append(messages, "user no created, check user payload")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
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
	var request UpdateUserRequest
	messages := make([]string, 0)

	// Parse the request body into the struct
	if err := c.BodyParser(&request); err != nil {
		messages = append(messages, "Error processing request")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	userID, err := strconv.Atoi(c.Params("id", "0"))

	if err != nil {
		messages = append(messages, "user id is not valid")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	user, err := h.UserService.GetUserByID(c.Context(), userID)

	if err != nil {
		messages = append(messages, "user not found")
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, messages)
	}

	userTypes, err := h.UserService.GetUserTypes(c.Context())

	if err != nil {
		messages = append(messages, "user types no available")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	if strings.TrimSpace(request.User.Name) != "" {
		user.Name = strings.TrimSpace(request.User.Name)
	}

	if strings.TrimSpace(request.User.Email) != "" {
		user.Name = strings.TrimSpace(request.User.Email)
	}

	if strings.TrimSpace(request.User.Password) != "" {
		password := strings.TrimSpace(request.User.Password)
		passwordHashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)

		if err != nil {
			messages = append(messages, "password is not valid")
			return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
		}

		user.Password = string(passwordHashed)
	}

	if strings.TrimSpace(request.User.Avatar) != "" {
		isImage, err := util.IsImageURL(strings.TrimSpace(request.User.Avatar))
		if err != nil {
			messages = append(messages, "avatar has to be a valid image")
		} else if !isImage {
			messages = append(messages, "avatar has to be a valid image")
		} else {
			user.Avatar = strings.TrimSpace(request.User.Avatar)
		}
	}

	if strings.TrimSpace(request.User.Phone) != "" {
		user.Phone = strings.TrimSpace(request.User.Phone)
	}

	var typesAvailable []int

	for _, userType := range userTypes {
		typesAvailable = append(typesAvailable, userType.ID)
	}

	if len(typesAvailable) > 0 && !util.Includes(typesAvailable, user.UserTypeID) {
		messages = append(messages, "user type id is not valid")
	}

	if len(messages) > 0 {
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	err = h.UserService.UpdateUser(c.Context(), &user)

	if err != nil {
		messages = append(messages, err.Error())
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	response := make(map[string]interface{})
	response["user"] = user

	// Return user as JSON
	return util.SuccessReponse(c, response)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	messages := make([]string, 0)
	userID, err := strconv.Atoi(c.Params("id", "0"))

	if err != nil {
		messages = append(messages, "user id is not valid")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	_, err = h.UserService.GetUserByID(c.Context(), userID)

	if err != nil {
		messages = append(messages, "user doesn't exist")
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, messages)
	}

	err = h.UserService.DeletedUser(c.Context(), userID)

	if err != nil {
		messages = append(messages, err.Error())
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	response := make(map[string]interface{})
	response["msg"] = fmt.Sprintf("user %d deleted", userID)

	// Return user as JSON
	return util.SuccessReponse(c, response)
}
