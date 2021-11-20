package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID string
	First_name string	
	Last_name string 	`gorm:"not null"`
	Email string		
}

func (u * User) BeforeCreate(tx *gorm.DB)  (err error) {
	u.ID = uuid.New().String()

	return
}