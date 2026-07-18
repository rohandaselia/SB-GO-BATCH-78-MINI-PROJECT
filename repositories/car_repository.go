package repositories

import (
	"database/sql"
	"rent-car-project/config"
	"rent-car-project/models"
)

func CreateCar(car models.Car) (int, error) {
	var carID int
	err := config.DB.QueryRow(
		`INSERT INTO cars (owner_id, brand, model, license_plate, description, price_per_day) 
		 VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		car.OwnerID, car.Brand, car.Model, car.LicensePlate, car.Description, car.PricePerDay,
	).Scan(&carID)
	return carID, err
}

func GetCarsByOwnerID(ownerID int) ([]models.Car, error) {
	rows, err := config.DB.Query(`SELECT id, brand, model, license_plate, description, price_per_day, status FROM cars WHERE owner_id = $1`, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []models.Car
	for rows.Next() {
		var car models.Car
		if err := rows.Scan(&car.ID, &car.Brand, &car.Model, &car.LicensePlate, &car.Description, &car.PricePerDay, &car.Status); err == nil {
			cars = append(cars, car)
		}
	}
	return cars, nil
}

func GetCarOwnerID(carID int) (int, error) {
	var ownerID int
	err := config.DB.QueryRow("SELECT owner_id FROM cars WHERE id = $1", carID).Scan(&ownerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return ownerID, nil
}

func UpdateCar(car models.Car) error {
	_, err := config.DB.Exec(
		`UPDATE cars SET brand = $1, model = $2, license_plate = $3, description = $4, price_per_day = $5 WHERE id = $6`,
		car.Brand, car.Model, car.LicensePlate, car.Description, car.PricePerDay, car.ID,
	)
	return err
}

func DeleteCar(carID, ownerID int) (int64, error) {
	result, err := config.DB.Exec("DELETE FROM cars WHERE id = $1 AND owner_id = $2", carID, ownerID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func SearchAvailableCars(startDateStr, endDateStr string) ([]models.Car, error) {
	query := `
		SELECT c.id, c.owner_id, c.brand, c.model, c.license_plate, c.description, c.price_per_day, c.status
		FROM cars c
		WHERE c.status = 'available' 
		  AND NOT EXISTS (
			  SELECT 1 
			  FROM bookings b 
			  WHERE b.car_id = c.id
				AND b.status IN ('pending', 'paid', 'active')
				AND b.start_date <= $2 
				AND b.end_date >= $1   
		  )`

	rows, err := config.DB.Query(query, startDateStr, endDateStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []models.Car
	for rows.Next() {
		var car models.Car
		if err := rows.Scan(&car.ID, &car.OwnerID, &car.Brand, &car.Model, &car.LicensePlate, &car.Description, &car.PricePerDay, &car.Status); err == nil {
			cars = append(cars, car)
		}
	}
	return cars, nil
}

func GetCarPriceAndStatus(carID int) (float64, string, error) {
	var price float64
	var status string
	err := config.DB.QueryRow("SELECT price_per_day, status FROM cars WHERE id = $1", carID).Scan(&price, &status)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, "", nil // not found
		}
		return 0, "", err
	}
	return price, status, nil
}
