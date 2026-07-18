package repositories

import "rent-car-project/config"

func InsertCarImage(carID int, imageURL string, isPrimary bool) (int, error) {
	var id int
	err := config.DB.QueryRow(
		`INSERT INTO car_images (car_id, image_url, is_primary) VALUES ($1, $2, $3) RETURNING id`,
		carID, imageURL, isPrimary,
	).Scan(&id)
	return id, err
}
