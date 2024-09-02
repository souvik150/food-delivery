package services

import (
	"github.com/google/uuid"

	"github.com/souvik150/food-delivery/internal/database"
	"github.com/souvik150/food-delivery/internal/models"
)

func CreateRestaurant(restaurant *models.Restaurant) error {
	return database.DB.Create(restaurant).Error
}

func UpdateRestaurant(id uuid.UUID, updatedData *models.Restaurant) (*models.Restaurant, error) {
	var restaurant models.Restaurant
	if err := database.DB.First(&restaurant, "id = ?", id).Error; err != nil {
		return nil, err
	}

	restaurant.Name = updatedData.Name
	restaurant.Address = updatedData.Address
	restaurant.Phone = updatedData.Phone
	restaurant.Email = updatedData.Email
	restaurant.OpeningTime = updatedData.OpeningTime
	restaurant.ClosingTime = updatedData.ClosingTime
	restaurant.Rating = updatedData.Rating

	if err := database.DB.Save(&restaurant).Error; err != nil {
		return nil, err
	}

	return &restaurant, nil
}

func DeleteRestaurant(id uuid.UUID) error {
	return database.DB.Delete(&models.Restaurant{}, "id = ?", id).Error
}

func GetRestaurantByID(id uuid.UUID) (*models.Restaurant, error) {
	var restaurant models.Restaurant
	if err := database.DB.First(&restaurant, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &restaurant, nil
}

func GetRestaurants(filter string, page int, limit int) ([]models.Restaurant, error) {
	var restaurants []models.Restaurant
	offset := (page - 1) * limit

	if err := database.DB.Offset(offset).Limit(limit).Find(&restaurants).Error; err != nil {
		return nil, err
	}

	return restaurants, nil
}
