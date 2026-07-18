package models

import "time"

type Car struct {
	ID           int         `json:"id"`
	OwnerID      int         `json:"owner_id"`
	Brand        string      `json:"brand"`
	Model        string      `json:"model"`
	LicensePlate string      `json:"license_plate"`
	Description  string      `json:"description"`
	PricePerDay  float64     `json:"price_per_day"`
	Status       string      `json:"status"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	Images       []CarImage  `json:"images,omitempty"`
}

type CarImage struct {
	ID        int       `json:"id"`
	CarID     int       `json:"car_id"`
	ImageURL  string    `json:"image_url"`
	IsPrimary bool      `json:"is_primary"`
	CreatedAt time.Time `json:"created_at"`
}

type CarInput struct {
	Brand        string  `json:"brand" binding:"required"`
	Model        string  `json:"model" binding:"required"`
	LicensePlate string  `json:"license_plate" binding:"required"`
	Description  string  `json:"description"`
	PricePerDay  float64 `json:"price_per_day" binding:"required,gt=0"`
}
