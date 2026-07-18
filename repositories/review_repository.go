package repositories

import "rent-car-project/config"

func InsertReview(bookingID, carID, customerID, rating int, comment string) (int, error) {
	var id int
	err := config.DB.QueryRow(
		`INSERT INTO reviews (booking_id, car_id, customer_id, rating, comment) 
		 VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		bookingID, carID, customerID, rating, comment,
	).Scan(&id)
	return id, err
}
