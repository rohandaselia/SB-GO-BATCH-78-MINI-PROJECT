package services

import (
	"errors"
	"rent-car-project/repositories"
	"strings"
)

func AddReview(customerID, bookingID, rating int, comment string) (int, error) {
	booking, err := repositories.GetBookingByID(bookingID)
	if err != nil {
		return 0, errors.New("booking not found")
	}

	if booking.CustomerID != customerID {
		return 0, errors.New("forbidden")
	}

	if booking.Status != "completed" {
		return 0, errors.New("booking must be completed to leave a review")
	}

	reviewID, err := repositories.InsertReview(bookingID, booking.CarID, customerID, rating, comment)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return 0, errors.New("review already exists for this booking")
		}
		return 0, errors.New("database error")
	}

	return reviewID, nil
}
