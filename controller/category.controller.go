package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/muazwzxv/bookers/m/service"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	gorm *gorm.DB
}

func NewCategoryRepository() *CategoryRepository {
	db := service.DB.DB
	return &CategoryRepository{gorm: db}
}

func (categoryRepo *CategoryRepository) Create(ctx *fiber.Ctx) error {

	// Response for created
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"Success":  true,
		"Message":  "Category Created",
		"Category": true,
	})
}

func (categoryRepo *CategoryRepository) Test(ctx *fiber.Ctx) error {
	fmt.Println("this is a test")
	return errors.New("this is a test")
}
