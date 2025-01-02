package main

import (
	"net/http"
	"github.com/blake-mcguire/golang-event-booking/main/models"
	"github.com/blake-mcguire/golang-event-booking/main/db"
	"github.com/gin-gonic/gin"
)


func main() {
	db.InitDB()
	server := gin.Default() //starts an instance of the engine 
	server.GET("/events", getEvents)//calling the GET function on the open instance of an engine in order to make get calls to the /events endpoint with the egt events fucntion passed
	server.POST("/events", createEvent)


	server.Run(":8080") //localhost:8080 its saying this server is running on this port 
}	

func getEvents(context *gin.Context) { //this context parameter will be set by gin and it will be a pointer to the gin context struct which contains a multidtude of object types
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events) 
}
	
func createEvent(context *gin.Context) {
	var event models.Event // setting the event variable as an struct of type event from the models package
	err := context.ShouldBindJSON(&event)	 //
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"mesage": "could not parse request data"})
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

	