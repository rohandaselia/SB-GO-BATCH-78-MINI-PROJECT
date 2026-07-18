package models

import "time"

type Payment struct {
	ID            int       `json:"id"`
	BookingID     int       `json:"booking_id"`
	Amount        float64   `json:"amount"`
	PaymentMethod string    `json:"payment_method"`
	PaymentStatus string    `json:"payment_status"`
	TransactionID string    `json:"transaction_id"`
	PaidAt        time.Time `json:"paid_at,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
}
