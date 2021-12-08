package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tag struct {
	ID                *string `json:"id"    gorm:"primarykey;type:varchar(256)"`
	Name              *string `json:"name"  gorm:"not null"`
	Is_system_defined *bool   `json:"-"     gorm:"index"`
	Tasks             *[]Tag  `json:"tasks" gorm:"many2many:task_tags;"`

	UserID *string `json:"-" gorm:"type:varchar;size:256"`
	User   *User
	BaseModel
}

type TaskTag struct {
	ID     *string `json:"id"    gorm:"primarykey;type:varchar(256)"`
	TagId  *string `gorm:"primarykey;type:varchar(256)`
	TaskId *string `gorm:"primarykey;type:varchar(256)`
	BaseModel
}

func (u *Tag) BeforeCreate(tx *gorm.DB) (err error) {

	newUUID := uuid.New().String()

	u.ID = &newUUID

	return
}

func (u *TaskTag) BeforeCreate(tx *gorm.DB) (err error) {

	newUUID := uuid.New().String()

	u.ID = &newUUID

	return
}
