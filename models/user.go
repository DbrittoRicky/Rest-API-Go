package models

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
