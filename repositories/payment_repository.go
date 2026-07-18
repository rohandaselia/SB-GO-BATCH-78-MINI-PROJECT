package repositories

import (
	"rent-car-project/config"
	"time"
)

func CreatePayment(bookingID int, amount float64) error {
	_, err := config.DB.Exec(
		`INSERT INTO payments (booking_id, amount, payment_status) VALUES ($1, $2, 'pending')`,
		bookingID, amount,
	)
	return err
}

func UpdatePaymentAndBookingStatus(bookingID int, status, transactionID, paymentMethod string, paidAt time.Time) error {
	tx, err := config.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	switch status {
	case "PAID", "SETTLED":
		_, err = tx.Exec(
			`UPDATE payments SET payment_status = 'paid', transaction_id = $1, payment_method = $2, paid_at = $3 WHERE booking_id = $4`,
			transactionID, paymentMethod, paidAt, bookingID,
		)
		if err != nil {
			return err
		}

		_, err = tx.Exec(`UPDATE bookings SET status = 'paid' WHERE id = $1`, bookingID)
		if err != nil {
			return err
		}
	case "EXPIRED":
		_, err = tx.Exec(`UPDATE payments SET payment_status = 'failed' WHERE booking_id = $1`, bookingID)
		if err != nil {
			return err
		}
		_, err = tx.Exec(`UPDATE bookings SET status = 'failed' WHERE id = $1`, bookingID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
