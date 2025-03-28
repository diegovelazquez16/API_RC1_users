package repository

import (
	"holamundo/users/domain/models"
)

type IUserRepository interface {
	Create(user *models.User) error
	GetLastFiveUsers() ([]models.User, error)
	CountByGender(gender string) (int64, error)
}

