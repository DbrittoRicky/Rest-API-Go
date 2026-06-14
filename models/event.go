// Package models
package models

import (
	"time"
)

type Event struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int       `gorm:"not null"`
	User        User      `json:"-" gorm:"foreignKey:UserID;references:ID" binding:"-"`
}
