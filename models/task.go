package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	BaseModel
	ID            *string    `json:"id"            gorm:"primarykey;type:varchar(256)"`
	Description   *string    `json:"description"   gorm:"not null"`
	Scheduled_for *time.Time `json:"scheduled_for" gorm:"index"`
	UserID        *string    `json:"user_id"       gorm:"type:varchar;size:256;not null"`
	User          *User
	Tags          *[]Tag `                     gorm:"many2many:task_tags;foreignKey:ID;joinForeignKey:TaskID;References:ID;joinReferences:TagID"`
	// TagNames      *[]string `json:"tags"`
}

func (u Task) String() string {
	return *u.Description
}

func (u *Task) BeforeCreate(tx *gorm.DB) (err error) {

	newUUID := uuid.New().String()

	u.ID = &newUUID
	// now := time.Now()

	// u.CreatedAt = now
	// u.UpdatedAt = &now
	return
}

// func (u *Task) AfterFind(tx *gorm.DB) (err error) {
// 	if u.Tags != nil {
// 		var tagNames []string

// 		for _, tag := range *u.Tags {
// 			tagNames = append(tagNames, *tag.Name)
// 		}

// 		u.TagNames = &tagNames
// 	}

// 	return
// }
