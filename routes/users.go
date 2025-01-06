package routes

import (
	"log"
	"net/http"
	"github.com/blake-mcguire/golang-event-booking/main/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("Error Binding JSON: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
	}

	err = user.Save()

	if err != nil {
		log.Printf("Error saving event: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "user created", "user": user })
}

func login(context *gin.Context) {
	var user models.User	
	err := context.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("Error Binding JSON: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return 
	}


	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"Message": err.Error()})
		return 
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login Successful"})
}	