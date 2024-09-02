package models

import (
    "gorm.io/gorm"

    "github.com/google/uuid"
)

type Cart struct {
    ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    CustomerID  uuid.UUID `gorm:"type:uuid;not null"`
    DishID      uuid.UUID `gorm:"type:uuid;not null"`
    Quantity    int       `gorm:"type:int;not null"`
    AddedDate   string    `gorm:"type:timestamp;not null"`

    Customer    Customer `gorm:"foreignKey:CustomerID"`
    Dish        Dish     `gorm:"foreignKey:DishID"`
}

func (c *Cart) BeforeCreate(tx *gorm.DB) (err error) {
    c.ID = uuid.New()
    return
}
