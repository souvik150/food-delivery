package schemas

import (
	"time"

	"github.com/google/uuid"
)

type CreateRestaurantSchema struct {
	Name         string    `json:"name" validate:"required,min=3,max=100"`
	Address      string    `json:"address" validate:"required,min=10,max=255"`
	Phone        string    `json:"phone" validate:"required,len=10"`
	Email        string    `json:"email" validate:"required,email"`
	OpeningTime  time.Time `json:"openingTime" validate:"required"`
	ClosingTime  time.Time `json:"closingTime" validate:"required"`
	Rating       float32   `json:"rating" validate:"required,min=0,max=5"`
	Country      string    `json:"country" validate:"required"`
	AverageSpend float64   `json:"averageSpend" validate:"required"`
	Cuisines     string    `json:"cuisines" validate:"required"`
}

type GetRestaurantByIDSchema struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Address      string    `json:"address"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}


type GetAllRestaurantsSchema struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Address      string    `json:"address"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	Rating       float64   `json:"rating"`
	Country      string    `json:"country"`
	AverageSpend float64   `json:"averageSpend"`
	Cuisines     string    `json:"cuisines"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}