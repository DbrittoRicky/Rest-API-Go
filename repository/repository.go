// Package repository
package repository

import (
	"errors"

	"rest-api.com/db"
	"rest-api.com/models"
	"rest-api.com/utils"
)

// ------------------------------- Event Operations -------------------------------------//

func GetEvents() ([]models.Event, error) {
	var events []models.Event

	allEvents := db.DB.Find(&events)

	return events, allEvents.Error
}

func GetEventByID(id string) (models.Event, error) {
	var event models.Event
	result := db.DB.First(&event, id)
	return event, result.Error
}

func CreateEvent(event *models.Event) {
	db.DB.Create(event)
}

func UpdateEvent(id string, updatedEvent *models.Event) (models.Event, error) {
	event, err := GetEventByID(id)
	if err != nil {
		return event, err
	}

	db.DB.Model(&event).Updates(updatedEvent)

	return event, nil
}

func RemoveEvent(e *models.Event) {
	db.DB.Delete(e)
}

// ----------------------------------------------- User Operations ---------------------------------//

func CreateUser(user *models.User) {
	db.DB.Create(user)
}

func ValidateUser(user models.User) (int, error) {
	var userRef models.User

	result := db.DB.Where("email = ?", user.Email).First(&userRef)

	if result.Error != nil {
		return userRef.ID, result.Error
	}

	retrievedPassword := userRef.Password

	isPasswordValid := utils.CheckPassword(user.Password, retrievedPassword)

	if !isPasswordValid {
		return userRef.ID, errors.New("invalid credentials")
	}

	return userRef.ID, nil
}
