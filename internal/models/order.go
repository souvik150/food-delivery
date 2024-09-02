package models

import (
    "gorm.io/gorm"

    "github.com/google/uuid"
)

type Order struct {
    ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    CustomerID   uuid.UUID `gorm:"type:uuid;not null"`
    RestaurantID uuid.UUID `gorm:"type:uuid;not null"`
    OrderDate    string    `gorm:"type:timestamp;not null"`
    TotalAmount  float64   `gorm:"type:decimal(10,2);not null"`
    Status       string    `gorm:"type:varchar(50);not null"`

    Customer     Customer   `gorm:"foreignKey:CustomerID"`
    Restaurant   Restaurant `gorm:"foreignKey:RestaurantID"`
    OrderItems   []OrderItem
    Deliveries   []Delivery
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
    o.ID = uuid.New()
    return
}
