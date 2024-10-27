package services

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/mhcodev/fake_store_api/internal/models"
	"github.com/mhcodev/fake_store_api/internal/repository/repositories"
	"github.com/mhcodev/fake_store_api/pkg"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepository repositories.AuthRepository
}

func NewAuthService(authRepository *repositories.AuthRepository) *AuthService {
	return &AuthService{
		authRepository: *authRepository,
	}
}

type LoginInput struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func (as *AuthService) Login(ctx context.Context, input LoginInput) (*models.User, error) {
	user, err := as.authRepository.GetUserByEmail(ctx, *input.Email)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(*input.Password))

	if err != nil {
		return nil, errors.New("email or password is not valid")
	}

	return &user, nil
}

func (as *AuthService) GetUserByToken(ctx context.Context, tokenStr string) (jwt.MapClaims, error) {

	claims, err := pkg.ExtractClaims(tokenStr)

	if err != nil {
		return nil, err
	}

	return claims, nil
}

type NewTokenInput struct {
	RefreshToken *string `json:"refreshToken"`
}

func (as *AuthService) GetNewToken(ctx context.Context, input NewTokenInput) (map[string]string, error) {
	claims, err := pkg.ExtractClaims(*input.RefreshToken)

	delete(claims, "exp")

	if err != nil {
		return nil, err
	}

	tokenMap := make(map[string]string, 0)

	accessToken, err := pkg.GenerateAccessToken(claims)

	if err != nil {
		return nil, err
	}

	refreshToken, err := pkg.GenerateRefreshToken(claims)

	tokenMap["accessToken"] = accessToken
	tokenMap["refreshToken"] = refreshToken

	if err != nil {
		return nil, err
	}

	return tokenMap, nil
}
