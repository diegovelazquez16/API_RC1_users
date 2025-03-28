package controllers

import (
	"net/http"
	"time"
	"holamundo/users/aplication/usecase"
	"github.com/gin-gonic/gin"
)

type UserGenderCountController struct {
	UserUC *usecase.GetGenderCountUseCase
}

func (uc *UserGenderCountController) GetGenderCount(ctx *gin.Context) {
	initialCounts, _ := uc.UserUC.Execute()

	for {
		time.Sleep(5 * time.Second) 

		newCounts, _ := uc.UserUC.Execute()
		if newCounts["masculino"] != initialCounts["masculino"] || newCounts["femenino"] != initialCounts["femenino"] {
			ctx.JSON(http.StatusOK, newCounts)
			return
		}
	}
}
