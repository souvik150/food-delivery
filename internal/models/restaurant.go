package models

import (
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

type Restaurant struct {
    ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    Name        string    `gorm:"type:varchar(255);not null"`
    Address     string    `gorm:"type:varchar(255);not null"`
    Phone       string    `gorm:"type:varchar(20);not null"`
    Email       string    `gorm:"type:varchar(255);not null"`
    OpeningTime  time.Time `gorm:"type:time"`
    ClosingTime  time.Time `gorm:"type:time"`
    Rating      float32   `gorm:"type:float"`
		Country     string    `gorm:"type:varchar(100);not null"`
		AvgSpendForTwo float32 `gorm:"type:float"`
		Cusines   string    `gorm:"type:varchar(255);not null"`

    Dishes      []Dish
    Orders      []Order
    Reviews     []Review
}

func (r *Restaurant) BeforeCreate(tx *gorm.DB) (err error) {
    r.ID = uuid.New()
    return
}
