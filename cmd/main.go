package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	config "github.com/souvik150/food-delivery/config"
	database "github.com/souvik150/food-delivery/internal/database"
	routes "github.com/souvik150/food-delivery/internal/routes"
)

func main() {
	app := fiber.New()

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}

	database.ConnectDB(&config)
	database.RunMigrations(database.DB)

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowCredentials: false,
	}))

	apiGroup := app.Group("/v1")

	routes.RegisterRestaurantRoutes(apiGroup)

	apiGroup.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Food Delivery App",
		})
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Route not found",
		})
	})

	log.Fatal(app.Listen(config.Port))
}