package controller

import (
	"example/restapi/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context){
	var input model.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err  !=nil{
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user := model.User{
		Username: input.Username,
		Password: input.Password,
	}
	savedUser,err := user.Save()

	if err !=nil{
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}
	context.JSON(http.StatusCreated,gin.H{"user":savedUser})
}