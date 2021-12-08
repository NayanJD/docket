package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID            *string    `json:"id"            gorm:"primarykey;type:varchar(256)"`
	Description   *string    `json:"description"   gorm:"not null"`
	Scheduled_for *time.Time `json:"scheduled_for" gorm:"index"`
	UserID        *string    `json:"user_id"       gorm:"type:varchar;size:256;not null"`
	User          *User
	Tags          *[]Tag `json:"tags"          gorm:"many2many:task_tags"`
	BaseModel
}

func (u *Task) BeforeCreate(tx *gorm.DB) (err error) {

	newUUID := uuid.New().String()

	u.ID = &newUUID

	return
}
