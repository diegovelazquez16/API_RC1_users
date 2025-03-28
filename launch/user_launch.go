package launch

import (
	"holamundo/core"
	userUsecase "holamundo/users/aplication/usecase"
	userRepo "holamundo/users/domain/repository"
	userControllers "holamundo/users/infraestructure/controllers"
	userRoutes "holamundo/users/infraestructure/routes"
	"github.com/gin-gonic/gin"
)

func RegisterUserModule(router *gin.Engine) {
	userRepo := &userRepo.UserRepositoryImpl{DB: core.GetDB()}

	createUserUC := &userUsecase.CreateUserUseCase{UserRepo: userRepo}
	lastFiveUC := &userUsecase.GetLastFiveUsersUseCase{UserRepo: userRepo}
	genderCountUC := &userUsecase.GetGenderCountUseCase{UserRepo: userRepo}

	userCreateController := &userControllers.UserCreateController{UserUC: createUserUC}
	userLastFiveController := &userControllers.UserLastFiveController{UserUC: lastFiveUC}
	userGenderCountController := &userControllers.UserGenderCountController{UserUC: genderCountUC}

	userRoutes.UserRoutes(router, userCreateController, userLastFiveController, userGenderCountController)
}
