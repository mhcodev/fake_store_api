package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mhcodev/fake_store_api/internal/services"
	"github.com/mhcodev/fake_store_api/internal/util"
	"github.com/mhcodev/fake_store_api/internal/validators"
	"github.com/mhcodev/fake_store_api/pkg"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var input services.LoginInput
	var validationErrors = validators.ValidationErrors{}

	if err := c.BodyParser(&input); err != nil {
		validationErrors.AddError("msg", "email & password are not valid")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	validationErrors = validators.ValidateLoginInput(input)

	if validationErrors.HasErrors() {
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	user, err := h.AuthService.Login(c.Context(), input)

	if err != nil {
		validationErrors.AddError("msg", err.Error())
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	data := make(map[string]interface{}, 0)
	data["user_id"] = user.ID

	accessToken, err := pkg.GenerateAccessToken(data)

	if err != nil {
		validationErrors.AddError("msg", "error generating access token")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	refreshToken, err := pkg.GenerateRefreshToken(data)

	if err != nil {
		validationErrors.AddError("msg", "error generating refresh token")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	response := make(map[string]interface{})
	response["user"] = user
	response["accessToken"] = accessToken
	response["refreshToken"] = refreshToken

	return util.SuccessReponse(c, response)
}

func (h *AuthHandler) GetTokenData(c *fiber.Ctx) error {
	var validationErrors = validators.ValidationErrors{}

	if len(c.GetReqHeaders()["Authorization"]) <= 0 {
		validationErrors.AddError("msg", "Token is not valid")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	bearerToken := c.GetReqHeaders()["Authorization"][0]

	if !strings.Contains(strings.ToLower(bearerToken), "bearer") {
		validationErrors.AddError("msg", "Token is not valid")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	tokenArr := strings.Split(bearerToken, "Bearer")
	tokenStr := strings.TrimSpace(tokenArr[1])

	data, err := h.AuthService.GetUserByToken(c.Context(), tokenStr)

	if err != nil {
		validationErrors.AddError("msg", err.Error())
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	response := make(map[string]interface{})
	response["data"] = data

	return util.SuccessReponse(c, response)
}

func (h *AuthHandler) AccessTokenFromRefreshToken(c *fiber.Ctx) error {
	var input services.NewTokenInput
	var validationErrors = validators.ValidationErrors{}

	if err := c.BodyParser(&input); err != nil {
		validationErrors.AddError("msg", "refreshToken is not valid")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	accessToken, err := h.AuthService.GetNewToken(c.Context(), input)

	if err != nil {
		validationErrors.AddError("msg", err.Error())
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	response := make(map[string]interface{})
	response["accessToken"] = accessToken

	return util.SuccessReponse(c, response)
}
