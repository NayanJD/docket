package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID            *string    `json:"id"`
	Description   *string    `json:"description"   gorm:"not null"`
	Scheduled_for *time.Time `json:"scheduled_for" gorm:"index"`
	UserID        *string    `json:"user_id"       gorm:"type:varchar;size:256"`
	User          *User
	BaseModel
}

func (u *Task) BeforeCreate(tx *gorm.DB) (err error) {

	newUUID := uuid.New().String()

	u.ID = &newUUID

	return
}
