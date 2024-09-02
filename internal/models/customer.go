package models

import (
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type Customer struct {
    ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    FirstName string    `gorm:"type:varchar(100);not null"`
    LastName  string    `gorm:"type:varchar(100);not null"`
    Email     string    `gorm:"type:varchar(255);not null;unique"`
    Phone     string    `gorm:"type:varchar(20);not null"`
    Address   string    `gorm:"type:varchar(255);not null"`
    CreatedAt string    `gorm:"type:timestamp;not null"`

    Orders    []Order
    Reviews   []Review
    CartItems []Cart
}

func (c *Customer) BeforeCreate(tx *gorm.DB) (err error) {
    c.ID = uuid.New()
    return
}
