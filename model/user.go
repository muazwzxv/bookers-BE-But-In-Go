package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint64 `gorm:"primary_key;auto_increment"`
	Name     string `gorm:"not null"`
	Email    string
	Role     string
	Password string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Listing []Listing
	Comment []Comment
}

func CreateUser(db *gorm.DB, user *User) (*User, error) {
	err := db.Debug().Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
