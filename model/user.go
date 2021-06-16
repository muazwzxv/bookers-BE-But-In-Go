package model

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
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

// Hash wrapper functions
func HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}
func CheckPasswordHash(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}

// Queries wrapper function
func CreateUser(db *gorm.DB, user *User) error {
	if err := db.Debug().Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUser(db *gorm.DB, email string) (User, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user); err != nil {
		return user, nil
	} else {
		return User{}, err.Error
	}
}

func CheckEmailExist(db *gorm.DB, email string) bool {
	var user User

	if res := db.Debug().Where("email = ?", email).First(&user); res != nil && res.RowsAffected == 0 {
		return false
	} else {
		fmt.Println(res.Error)
		return true
	}
}
