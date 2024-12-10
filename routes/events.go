package routes

import (
	"net/http"
	"registrationApp/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve all the data"})
	}
	context.JSON(200, events)
}

func getEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse id."})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not fetch the id", "error": event})
		return
	}
	context.JSON(200, event)
}

func CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		// Include error details to provide better feedback to the client
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the object due to missing or invalid fields",
			"error":   err.Error(),
		})
		return
	}

	userId := context.GetInt64("userId")

	// Set default values for the event if needed
	event.UserID = userId

	err = event.Save()
	if err != nil {
		// Respond with a more appropriate status code for database-related errors
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event. Try again!",
			"error":   err.Error(),
		})
		return
	}

	// Successful response with 201 status code (Created)
	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created!",
		"event":   event,
	})
}

func UpdateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the id"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Could not find the id"})
		return
	}

	if userId != event.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update this event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBind(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the event ", "error": err})
		return
	}
	context.JSON(200, gin.H{"message": "event updated successfully"})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find the id"})
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the data"})
	}
	if userId != event.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete", "error": err})
	}

	context.JSON(200, gin.H{"message": "Event Deleted"})
}
