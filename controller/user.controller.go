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

	// Parse incoming request
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Success": false,
			"Message": "Cannot parse JSON",
		})
	}

	// Check email exits
	if cond := model.CheckEmailExist(userRepo.gorm, user.Email); cond == true {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Success": false,
			"Message": "Email already exists",
		})
	}

	// Hash password
	if hashed, err := model.HashPassword(user.Password); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Success": false,
			"Message": "Failed to hash password",
		})
	} else {
		user.Password = hashed
	}

	// Create the user
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

func (userRepo *UserRepository) Login(ctx *fiber.Ctx) error {
	// Parse request
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

	// Check password hash
	if model.CheckPasswordHash(login.Password, user.Password) == false {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Success": false,
			"Message": "Password does not match",
		})
	}

	// Generate jwt token
	if token, err := model.CreateToken(uint(user.ID)); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Success": false,
			"Message": err,
		})
	} else {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"Success": true,
			"token":   token,
		})
	}

}
