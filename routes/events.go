package routes

import (
	"log"
	"net/http"
	"github.com/blake-mcguire/golang-event-booking/main/models"

	"github.com/gin-gonic/gin"
	"strconv"
)


func getEvents(context *gin.Context) { //this context parameter will be set by gin and it will be a pointer to the gin context struct which contains a multidtude of object types
	events, err := models.GetAllEvents()

	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events!"})
		return
	}
	context.JSON(http.StatusOK, events) 
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) //this function allows you to pull the parameter from the call and extract the id the user is searching for
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event id "})
		return
	}

	event, err := models.GetEventById(eventId)
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"mesage":"Could not get the event!"})
		return
	}

	context.JSON(http.StatusOK, event)


}
	
func createEvent(context *gin.Context) {
	var event models.Event // setting the event variable as an struct of type event from the models package
	err := context.ShouldBindJSON(&event)	 //
	if err != nil {
		log.Printf("Error binding JSON: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"mesage": "could not parse request data"})
	}

	// event.ID = 1
	// event.UserID = 1

	err = event.Save()

	if err != nil {
		log.Printf("Error saving event: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"mesage": "could not parse request data"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}


func UpdateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event id "})
		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not fetch the proper path"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not fetch the proper path"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event successfully updated"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event id "})
		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot delete an event that does not exist!"})
		return
	}
	var deadEvent models.Event
	deadEvent.ID = eventId
	err = deadEvent.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not delete"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event successfully deleted"})
}