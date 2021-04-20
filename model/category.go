package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID   uint64 `gorm:"primary_key;auto_increment"`
	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	Topic []Topic
}