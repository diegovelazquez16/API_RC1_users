package controllers

import (
	"net/http"
	"holamundo/users/aplication/usecase"
	"holamundo/users/domain/models"
	"github.com/gin-gonic/gin"
)

type UserCreateController struct {
	UserUC *usecase.CreateUserUseCase
}

func (uc *UserCreateController) Create(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.UserUC.UserRepo.Create(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear usuario"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Usuario creado"})
}
