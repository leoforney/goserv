package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `json:"username" gorm:"unique;not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `gorm:"not null"`
}

func CreateUser(db *gorm.DB, user User) error {
	result := db.Create(&user)
	return result.Error
}

func GetUserByUsername(db *gorm.DB, username string, user *User) error {
	result := db.Where("username = ?", username).First(user)
	return result.Error
}
