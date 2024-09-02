package models

import (
    "gorm.io/gorm"

    "github.com/google/uuid"
)

type Dish struct {
    ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    Name         string    `gorm:"type:varchar(255);not null"`
    Description  string    `gorm:"type:text;not null"`
    Price        float64   `gorm:"type:decimal(10,2);not null"`
    RestaurantID uuid.UUID `gorm:"type:uuid;not null"`
    Category     string    `gorm:"type:varchar(100);not null"`

    Restaurant   Restaurant `gorm:"foreignKey:RestaurantID"`
    OrderItems   []OrderItem
    CartItems    []Cart
}

func (d *Dish) BeforeCreate(tx *gorm.DB) (err error) {
    d.ID = uuid.New()
    return
}
