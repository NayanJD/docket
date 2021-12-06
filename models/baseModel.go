package models

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index"    json:"deleted_at"`
}
