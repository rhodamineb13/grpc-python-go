package entities

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	gorm.Model
	Name        string    `gorm:"column:name"`
	Email       string    `gorm:"column:email;unique"`
	Username    string    `gorm:"column:username;unique"`
	Password    string    `gorm:"column:password"`
	DateOfBirth time.Time `gorm:"column:date_of_birth"`
}
