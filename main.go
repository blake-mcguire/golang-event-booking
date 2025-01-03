package main

import (
	// "log"
	// "net/http"
	// "strconv"

	"github.com/blake-mcguire/golang-event-booking/main/db"
	// "github.com/blake-mcguire/golang-event-booking/main/models"
	"github.com/blake-mcguire/golang-event-booking/main/routes"
	"github.com/gin-gonic/gin"
)


func main() {
	db.InitDB()
	server := gin.Default() //starts an instance of the engine 
	routes.RegisterRoutes(server)
	server.Run(":8080") //localhost:8080 its saying this server is running on this port

}	



	