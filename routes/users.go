package routes

import (
	"net/http"

	"exmaple.com/event-booking-rest-api/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user", "error": err})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
