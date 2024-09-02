package models

import (
    "gorm.io/gorm"

    "github.com/google/uuid"
)

type OrderItem struct {
    ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    OrderID  uuid.UUID `gorm:"type:uuid;not null"`
    DishID   uuid.UUID `gorm:"type:uuid;not null"`
    Quantity int       `gorm:"type:int;not null"`
    Price    float64   `gorm:"type:decimal(10,2);not null"`

    Order    Order `gorm:"foreignKey:OrderID"`
    Dish     Dish  `gorm:"foreignKey:DishID"`
}

func (oi *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
    oi.ID = uuid.New()
    return
}
