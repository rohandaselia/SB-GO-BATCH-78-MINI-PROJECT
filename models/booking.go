package models

import "time"

type Booking struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	CarID      int       `json:"car_id"`
	StartDate  string    `json:"start_date"` // Menggunakan string untuk format YYYY-MM-DD atau time.Time
	EndDate    string    `json:"end_date"`
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type BookingInput struct {
	CarID     int    `json:"car_id" binding:"required"`
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
}

type BookingStatusInput struct {
	Status string `json:"status" binding:"required,oneof=active completed"`
}
