package model

import (
	"time"

	"gorm.io/gorm"
)

type Listing struct {
	ID          uint64 `gorm:"primary_key;auto_increment"`
	Title       string
	Delivery    string
	Contact     string
	Description string
	price       float64

	//Foreign key
	UserId uint64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
