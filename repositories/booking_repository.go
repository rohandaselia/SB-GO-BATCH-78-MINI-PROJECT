package repositories

import (
	"rent-car-project/config"
	"rent-car-project/models"
)

func CreateBooking(customerID, carID int, startDate, endDate string, totalPrice float64) (int, error) {
	var bookingID int
	err := config.DB.QueryRow(
		`INSERT INTO bookings (customer_id, car_id, start_date, end_date, total_price) 
		 VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		customerID, carID, startDate, endDate, totalPrice,
	).Scan(&bookingID)
	return bookingID, err
}

func GetBookingByID(bookingID int) (*models.Booking, error) {
	var booking models.Booking
	err := config.DB.QueryRow(
		`SELECT id, customer_id, car_id, start_date, end_date, total_price, status 
		 FROM bookings WHERE id = $1`, bookingID,
	).Scan(&booking.ID, &booking.CustomerID, &booking.CarID, &booking.StartDate, &booking.EndDate, &booking.TotalPrice, &booking.Status)

	if err != nil {
		return nil, err
	}
	return &booking, nil
}

func UpdateBookingStatus(bookingID int, status string) error {
	_, err := config.DB.Exec(`UPDATE bookings SET status = $1 WHERE id = $2`, status, bookingID)
	return err
}
