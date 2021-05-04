package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/muazwzxv/bookers/m/model"
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

	var category model.Category

	if err := ctx.BodyParser(&category); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Success": false,
			"Message": "Cannot parse JSON",
		})
	}

	if err := model.CreateCategory(categoryRepo.gorm, &category); err != nil {
		return ctx.Status(http.StatusConflict).JSON(fiber.Map{
			"Success": false,
			"Message": "Cannot create user",
		})
	} else {
		// Response for created
		return ctx.Status(http.StatusCreated).JSON(fiber.Map{
			"Success":  true,
			"Message":  "Category Created",
			"Category": category,
		})
	}
}

func (categoryRepo *CategoryRepository) Test(ctx *fiber.Ctx) error {
	fmt.Println("this is a test")
	return errors.New("this is a test")
}
