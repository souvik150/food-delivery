package routes

import (
	"github.com/gofiber/fiber/v2"

	controllers "github.com/souvik150/food-delivery/internal/controllers"
)

func RegisterRestaurantRoutes(group fiber.Router) {
	api := group.Group("/restaurants")

	api.Post("/", controllers.CreateRestaurant)
	api.Get("/", controllers.GetRestaurants)
	api.Get("/:id", controllers.GetRestaurantByID)
	api.Put("/:id", controllers.UpdateRestaurant)
}