package services

import (
	"errors"
	"rent-car-project/repositories"
	"strings"
	"time"
)

func CreateBooking(customerID, carID int, startDateStr, endDateStr string) (int, float64, string, error) {
	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, startDateStr) // Already validated in controller
	endDate, _ := time.Parse(layout, endDateStr)

	days := int(endDate.Sub(startDate).Hours()/24) + 1

	pricePerDay, status, err := repositories.GetCarPriceAndStatus(carID)
	if err != nil {
		return 0, 0, "", errors.New("database error")
	}
	if pricePerDay == 0 {
		return 0, 0, "", errors.New("car not found")
	}
	if status != "available" {
		return 0, 0, "", errors.New("car is not available")
	}

	totalPrice := float64(days) * pricePerDay

	bookingID, err := repositories.CreateBooking(customerID, carID, startDateStr, endDateStr, totalPrice)
	if err != nil {
		if strings.Contains(err.Error(), "prevent_double_booking") {
			return 0, 0, "", errors.New("double_booking")
		}
		return 0, 0, "", err
	}

	invoiceURL, err := CreateInvoice(bookingID, totalPrice)
	if err != nil {
		return bookingID, totalPrice, "", errors.New("payment gateway error: " + err.Error())
	}

	return bookingID, totalPrice, invoiceURL, nil
}

func UpdateBookingStatusByOwner(ownerID, bookingID int, newStatus string) error {
	booking, err := repositories.GetBookingByID(bookingID)
	if err != nil {
		return errors.New("booking not found")
	}

	carOwnerID, err := repositories.GetCarOwnerID(booking.CarID)
	if err != nil || carOwnerID != ownerID {
		return errors.New("forbidden")
	}

	if newStatus == "active" && booking.Status != "paid" {
		return errors.New("invalid status transition: must be paid to become active")
	}
	if newStatus == "completed" && booking.Status != "active" {
		return errors.New("invalid status transition: must be active to become completed")
	}

	return repositories.UpdateBookingStatus(bookingID, newStatus)
}
