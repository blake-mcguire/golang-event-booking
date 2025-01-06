package routes

import (
	"log"
	"net/http"

	"github.com/blake-mcguire/golang-event-booking/main/models"
	"github.com/blake-mcguire/golang-event-booking/main/utils"
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

	token, err:= utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token!", "error": err})
		return
	}


	context.JSON(http.StatusOK, gin.H{"message": "Login Successful","token":token})
}	