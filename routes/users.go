package routes

import (
	"net/http"

	"github.com/chrisjoyce54/GoApi/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request date: " + err.Error() + "."})
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not save user: " + err.Error() + "."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created."})
}
