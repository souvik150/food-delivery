package controllers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/souvik150/food-delivery/internal/database"
	"github.com/souvik150/food-delivery/internal/models"
	"github.com/souvik150/food-delivery/internal/services"
)

func CreateRestaurant(c *fiber.Ctx) error {
	restaurant := new(models.Restaurant)
	if err := c.BodyParser(restaurant); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := services.CreateRestaurant(restaurant); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create restaurant"})
	}

	return c.Status(http.StatusCreated).JSON(restaurant)
}

func UpdateRestaurant(c *fiber.Ctx) error {
	id := c.Params("id")
	restaurant := new(models.Restaurant)
	if err := c.BodyParser(restaurant); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	updatedRestaurant, err := services.UpdateRestaurant(uuid.MustParse(id), restaurant)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Could not update restaurant"})
	}

	return c.Status(http.StatusOK).JSON(updatedRestaurant)
}

func DeleteRestaurant(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := services.DeleteRestaurant(uuid.MustParse(id)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete restaurant"})
	}
	return c.Status(http.StatusNoContent).Send(nil)
}

func GetRestaurantByID(c *fiber.Ctx) error {
	id := c.Params("id")
	restaurant, err := services.GetRestaurantByID(uuid.MustParse(id))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Restaurant not found"})
	}
	return c.Status(http.StatusOK).JSON(restaurant)
}

func GetRestaurants(c *fiber.Ctx) error {
	var restaurants []models.Restaurant
	
	// Convert query parameters to integers
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		limit = 10
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	filter := c.Query("filter", "")
	query := database.DB.Limit(limit).Offset(offset)
	if filter != "" {
		query = query.Where("name ILIKE ? ", "%"+filter+"%")
	}

	if err := query.Find(&restaurants).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch restaurants"})
	}

	return c.Status(http.StatusOK).JSON(restaurants)
}
