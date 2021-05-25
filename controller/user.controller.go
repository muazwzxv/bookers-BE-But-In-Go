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

func NewUserRepository() *UserRepository {
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

	if err := model.CreateUser(userRepo.gorm, &user); err != nil {
		return ctx.Status(http.StatusConflict).JSON(fiber.Map{
			"Success": false,
			"Message": "Cannot create user",
		})
	} else {
		return ctx.Status(http.StatusCreated).JSON(fiber.Map{
			"Success": true,
			"Message": "User Created",
			"User":    user,
		})
	}
}

func (userRepo *UserRepository) login(ctx *fiber.Ctx) error {
	var login model.Login
	if err := ctx.BodyParser(&login); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Success": false,
			"Message": "Cannot parse JSON",
		})
	}

	// check if user exist or not
	user, err := model.GetUser(userRepo.gorm, login.Email)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Success": false,
			"Message": err,
		})
	}
}
