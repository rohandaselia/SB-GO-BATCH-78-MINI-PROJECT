package models

import "time"

type Review struct {
	ID         int       `json:"id"`
	BookingID  int       `json:"booking_id"`
	CarID      int       `json:"car_id"`
	CustomerID int       `json:"customer_id"`
	Rating     int       `json:"rating"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at"`
}

type ReviewInput struct {
	Rating  int    `json:"rating" binding:"required,min=1,max=5"`
	Comment string `json:"comment"`
}
