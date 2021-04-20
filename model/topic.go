package model

import (
	"time"

	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	ID   uint64 `gorm:"primary_key;auto_increment"`
	Name string

	CategoryId uint64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	Comment []Comment
}
