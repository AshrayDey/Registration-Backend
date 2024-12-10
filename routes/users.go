package routes

import (
	"net/http"
	"registrationApp/models"
	"registrationApp/utils"

	"github.com/gin-gonic/gin"
)

// func getAllEvents(context *gin.Context) {

// }

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Missing fields", "error": err})
		return
	}
	err = user.Save()
	context.JSON(http.StatusOK, gin.H{"message": "User created"})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateJWT(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Token in valid"})
		return
	}
	context.JSON(200, gin.H{"message": "Logged in!!", "token": token})
}

func getUsers(context *gin.Context) {

	users, err := models.GetAllUser()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not retrieve data from user table"})
	}
	context.JSON(200, users)
}
