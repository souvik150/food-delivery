package models

import (
    "gorm.io/gorm"

    "github.com/google/uuid"
)

type Delivery struct {
    ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    OrderID       uuid.UUID `gorm:"type:uuid;not null"`
    DeliveryPerson string    `gorm:"type:varchar(100);not null"`
    DeliveryTime  string    `gorm:"type:timestamp;not null"`
    Status        string    `gorm:"type:varchar(50);not null"`

    Order         Order `gorm:"foreignKey:OrderID"`
}

func (d *Delivery) BeforeCreate(tx *gorm.DB) (err error) {
    d.ID = uuid.New()
    return
}
