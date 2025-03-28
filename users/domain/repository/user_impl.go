package repository

import (
	"gorm.io/gorm"
	"holamundo/users/domain/models"
)


type UserRepositoryImpl struct {
	DB *gorm.DB  

}


func (r *UserRepositoryImpl) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepositoryImpl) GetLastFiveUsers() ([]models.User, error) {
	var users []models.User
	err := r.DB.Order("id DESC").Limit(5).Find(&users).Error
	return users, err
}

func (r *UserRepositoryImpl) CountByGender(gender string) (int64, error) {
	var count int64
	err := r.DB.Model(&models.User{}).Where("genero = ?", gender).Count(&count).Error
	return count, err
}
