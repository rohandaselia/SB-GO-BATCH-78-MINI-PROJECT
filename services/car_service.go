package services

import (
	"errors"
	"rent-car-project/models"
	"rent-car-project/repositories"
)

func CreateCar(ownerID int, brand, model, licensePlate, description string, price float64) (int, error) {
	car := models.Car{
		OwnerID:      ownerID,
		Brand:        brand,
		Model:        model,
		LicensePlate: licensePlate,
		Description:  description,
		PricePerDay:  price,
	}
	return repositories.CreateCar(car)
}

func GetMyCars(ownerID int) ([]models.Car, error) {
	return repositories.GetCarsByOwnerID(ownerID)
}

func UpdateCar(ownerID, carID int, brand, model, licensePlate, description string, price float64) error {
	dbOwnerID, err := repositories.GetCarOwnerID(carID)
	if err != nil {
		return errors.New("database error")
	}
	if dbOwnerID == 0 {
		return errors.New("car not found")
	}
	if dbOwnerID != ownerID {
		return errors.New("forbidden")
	}

	car := models.Car{
		ID:           carID,
		Brand:        brand,
		Model:        model,
		LicensePlate: licensePlate,
		Description:  description,
		PricePerDay:  price,
	}
	return repositories.UpdateCar(car)
}

func DeleteCar(ownerID, carID int) error {
	rows, err := repositories.DeleteCar(carID, ownerID)
	if err != nil {
		return errors.New("failed to delete car")
	}
	if rows == 0 {
		return errors.New("not found or forbidden")
	}
	return nil
}

func SearchCars(startDateStr, endDateStr string) ([]models.Car, error) {
	return repositories.SearchAvailableCars(startDateStr, endDateStr)
}
