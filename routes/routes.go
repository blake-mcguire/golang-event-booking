package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine){
	server.GET("/events", getEvents)//calling the GET function on the open instance of an engine in order to make get calls to the /events endpoint with the egt events fucntion passed
	server.GET("/events/:id", getEvent) // this is how to set up a slug to attrubute an id in gin to get a event by its id
	server.POST("/events", createEvent)
	server.PUT("/events/:id", UpdateEvent)	
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup)
}