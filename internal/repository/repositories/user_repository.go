package repositories

import "github.com/mhcodev/fake_store_api/internal/models"

type UserRepository interface {
	GetUsersByParams() ([]models.User, error)
}
