// Package handler
package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"rest-api.com/models"
	"rest-api.com/repository"
	"rest-api.com/utils"
)

//------------------------------------------------------- Event handlers -------------------------------------------//

func GetEvents(context *gin.Context) {
	events, err := repository.GetEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, events)
}

func GetEvent(context *gin.Context) {
	event, err := repository.GetEventByID(context.Param("id"))

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func CreateEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetInt("userId")

	event.UserID = userId

	repository.CreateEvent(&event)

	context.JSON(http.StatusCreated, event)
}

func UpdateEvent(context *gin.Context) {
	var updated models.Event

	userId := context.GetInt("userId")

	if updated.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}

	err := context.ShouldBindJSON(&updated)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := repository.UpdateEvent(context.Param("id"), &updated)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, event)
}

func RemoveEvent(context *gin.Context) {
	event, err := repository.GetEventByID(context.Param("id"))

	userId := context.GetInt("userId")

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	repository.RemoveEvent(&event)

	context.JSON(http.StatusOK, gin.H{"Alert": "event removed"})
}

//----------------------------------------------------------------------- User handlers ------------------------------------------//

func CreateUser(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password := user.Password

	hashedPassword := utils.HashPassword(password)

	user.Password = hashedPassword

	repository.CreateUser(&user)

	context.JSON(http.StatusCreated, user)
}

func Login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := repository.ValidateUser(user)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	token, err := utils.GenerateToken(user.Email, userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not authenticate user", "id": userId})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login Success", "token": token})
}
