package routes

import (
	"net/http"
	"strconv"

	"github.com/chrisjoyce54/GoApi/models"
	"github.com/chrisjoyce54/GoApi/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request date: " + err.Error() + "."})
	}

	u, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not save user: " + err.Error() + "."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created: " + strconv.FormatInt(u.ID, 10) + " " + u.Email})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request: " + err.Error()})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate: " + err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Success", "token": token})
}

func getUsers(context *gin.Context) {
	events, err := models.GetUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch users. Try again later: " + err.Error() + "."})
		return
	}
	context.JSON(http.StatusOK, events)
}
