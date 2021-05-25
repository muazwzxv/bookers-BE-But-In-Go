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

type Login struct {
	Email    string
	Password string
}

func CreateUser(db *gorm.DB, user *User) error {
	if err := db.Debug().Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUser(db *gorm.DB, email string) (User, error) {
	var user User
	if err := db.Debug().Where("email = ?", email).First(&user); err != nil {
		return user, nil
	} else {
		return User{}, err.Error
	}
}

func CheckEmailExist(db *gorm.DB, email string) bool {
	var user User
	if err := db.Debug().Where("email = ?", email).First(&user); err != nil {
		return true
	} else {
		return false
	}
}
