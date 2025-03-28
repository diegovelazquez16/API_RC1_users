package usecase

import (
	"holamundo/users/domain/models"
	"holamundo/users/domain/repository"
)

type CreateUserUseCase struct {
	UserRepo repository.IUserRepository
}

func (uc *CreateUserUseCase) Execute(user *models.User) error {
	return uc.UserRepo.Create(user)
}

type GetLastFiveUsersUseCase struct {
	UserRepo repository.IUserRepository
}

func (uc *GetLastFiveUsersUseCase) Execute() ([]models.User, error) {
	return uc.UserRepo.GetLastFiveUsers()
}

type GetGenderCountUseCase struct {
	UserRepo repository.IUserRepository
}

func (uc *GetGenderCountUseCase) Execute() (map[string]int, error) {
	genderCounts := make(map[string]int)

	maleCount, err := uc.UserRepo.CountByGender("masculino")
	if err != nil {
		return nil, err
	}

	femaleCount, err := uc.UserRepo.CountByGender("femenino")
	if err != nil {
		return nil, err
	}

	genderCounts["masculino"] = int(maleCount)
	genderCounts["femenino"] = int(femaleCount)

	return genderCounts, nil
}

