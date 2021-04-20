package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/muazwzxv/bookers/m/model"
	"github.com/muazwzxv/bookers/m/service"
	"gorm.io/gorm"
)

type UserRepository struct {
	gorm *gorm.DB
}

func New() *UserRepository {
	db := service.DB.DB
	return &UserRepository{gorm: db}
}

func (userRepo *UserRepository) Create(ctx *fiber.Ctx) error {
	var user model.User

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Success": false,
			"Message": "Cannot parse JSON",
		})
	}

	created, err := model.CreateUser(userRepo.gorm, &user)
	if err != nil {
		return ctx.Status(http.StatusConflict).JSON(fiber.Map{
			"Success": false,
			"Message": "Cannot create user",
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"Success": true,
		"Message": "User Created",
		"User":    created,
	})
}
