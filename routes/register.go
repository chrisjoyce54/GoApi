package routes

import (
	"net/http"
	"strconv"

	"github.com/chrisjoyce54/GoApi/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id: " + err.Error() + "."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event: " + err.Error()})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could register user for event: " + err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Registered"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id: " + err.Error() + "."})
		return
	}
	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not cancel registratioin: " + err.Error() + "."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Cancelled"})
}
