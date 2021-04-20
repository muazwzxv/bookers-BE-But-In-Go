package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID          uint64 `gorm:"primary_key;auto_increment"`
	Description string

	TopicID uint64
	UserId  uint64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
