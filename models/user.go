package models

import (
	"fmt"
	"net/mail"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	BaseModel
	ID           *string `json:"id"			gorm:"primarykey;type:varchar;size:256"`
	First_name   *string `json:"first_name" gorm:"not null"`
	Last_name    *string `json:"last_name"  gorm:"not null"`
	Username     *string `json:"username"   gorm:"not null;unique"`
	Password     *string `json:"-"          gorm:"not null"`
	Is_superuser *bool   `json:"-"          gorm:"not null;default:false"`
	Is_staff     *bool   `json:"-"          gorm:"not null;default:false"`
	Tasks        *[]Task `gorm:"type:Task[];foreignKey:UserID;references:ID"`
}

func (u *User) String() string {
	return fmt.Sprintf("User: id=&v, username=%v", *u.ID, *u.Username)
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	newUUID := uuid.New().String()

	u.ID = &newUUID

	if u.Username != nil {
		_, err = mail.ParseAddress(*u.Username)

		if err != nil {
			return err
		}
	}

	if u.Password != nil {
		hash, err := bcrypt.GenerateFromPassword([]byte(*u.Password), bcrypt.MinCost)

		if err != nil {
			return err
		}

		*u.Password = string(hash)
	}

	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	*u.ID = uuid.New().String()

	if tx.Statement.Changed("Username") {
		_, err = mail.ParseAddress(*u.Username)

		if err != nil {
			return err
		}
	}

	if tx.Statement.Changed("Password") {
		hash, err := bcrypt.GenerateFromPassword([]byte(*u.Password), bcrypt.MinCost)

		if err != nil {
			return err
		}

		*u.Password = string(hash)
	}

	return
}

func (u *User) ComparePassword(pwd *string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(*u.Password), []byte(*pwd))

	if err != nil {
		log.Error().Err(err)
		return false
	}

	return true
}
