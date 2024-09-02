package database

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	models "github.com/souvik150/food-delivery/internal/models"
)

func RunMigrations(db *gorm.DB) {
	log.Println("Running Migrations")

	err := db.AutoMigrate(
		&models.Restaurant{},
		&models.Dish{},
		&models.Order{},
		&models.OrderItem{},
		&models.Review{},
		&models.Customer{},
		&models.Delivery{},
		&models.Cart{},
	)

	if err != nil {
		fmt.Println("Migration error")
		return
	}

	log.Println("🚀 Migrations completed")
}