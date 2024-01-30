package controller

import (
	"example/restapi/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)
func GetAllEntries(context *gin.Context){
	user, err := helper.CurrentUser(context)
	if err !=nil{
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	context.JSON(http.StatusOK,gin.H{"data":user.Entries})
}