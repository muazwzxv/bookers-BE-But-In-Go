package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/muazwzxv/bookers/m/controller"
	"github.com/muazwzxv/bookers/m/service"
)

func main() {
	//app := fiber.New()

	// Connect Database
	if _, err := service.DB.Connect(); err != nil {
		log.Fatal("Error ", err)
	}

	app := setupRouter()
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint",
		})
	})

	err := app.Listen(":8000")
	if err != nil {
		log.Fatal(err)
	}
}

func setupRouter() *fiber.App {
	app := fiber.New()

	userRepository := controller.NewUserRepository()
	app.Post("/users", userRepository.Create)

	categoryRepository := controller.NewCategoryRepository()
	app.Get("/test", categoryRepository.Test)
	app.Post("/category", categoryRepository.Create)

	return app
}
