package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID string
	First_name string	`gorm:"not null"`
	Last_name string 	`gorm:"not null"`
	Email string		`gorm:"unique"`
}

func (u * User) BeforeCreate(tx *gorm.DB)  (err error) {
	u.ID = uuid.New().String()

	return
}