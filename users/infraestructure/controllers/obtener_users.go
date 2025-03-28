package controllers

import (
	"net/http"
	"holamundo/users/aplication/usecase"
	"github.com/gin-gonic/gin"
)

type UserLastFiveController struct {
	UserUC *usecase.GetLastFiveUsersUseCase
}

func (uc *UserLastFiveController) GetLastFiveUsers(ctx *gin.Context) {
	users, err := uc.UserUC.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios"})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
