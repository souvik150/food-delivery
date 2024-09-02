package models

import (
    "gorm.io/gorm"

    "github.com/google/uuid"
)

type Review struct {
    ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    CustomerID   uuid.UUID `gorm:"type:uuid;not null"`
    RestaurantID uuid.UUID `gorm:"type:uuid;not null"`
    Rating       int       `gorm:"type:int;not null"`
    Comment      string    `gorm:"type:text"`
    ReviewDate   string    `gorm:"type:timestamp;not null"`

    Customer     Customer   `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
    Restaurant   Restaurant `gorm:"foreignKey:RestaurantID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (r *Review) BeforeCreate(tx *gorm.DB) (err error) {
    r.ID = uuid.New()
    return
}
