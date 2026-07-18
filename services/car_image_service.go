package services

import (
	"errors"
	"rent-car-project/repositories"
)

func AddCarImage(ownerID, carID int, imageURL string, isPrimary bool) (int, error) {
	dbOwnerID, err := repositories.GetCarOwnerID(carID)
	if err != nil {
		return 0, errors.New("database error")
	}
	if dbOwnerID == 0 {
		return 0, errors.New("car not found")
	}
	if dbOwnerID != ownerID {
		return 0, errors.New("forbidden")
	}

	return repositories.InsertCarImage(carID, imageURL, isPrimary)
}
