package controllers

import (
	"fmt"
	"rent-car-project/models"
	"net/http"
	"rent-car-project/services"
	"time"

	"github.com/gin-gonic/gin"
)


func CreateBooking(c *gin.Context) {
	customerID := c.GetInt("user_id")

	var input models.BookingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	layout := "2006-01-02"
	startDate, err := time.Parse(layout, input.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format, use YYYY-MM-DD"})
		return
	}
	endDate, err := time.Parse(layout, input.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format, use YYYY-MM-DD"})
		return
	}

	if endDate.Before(startDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "end_date cannot be before start_date"})
		return
	}

	bookingID, totalPrice, invoiceURL, err := services.CreateBooking(customerID, input.CarID, input.StartDate, input.EndDate)
	if err != nil {
		switch err.Error() {
		case "car not found":
			c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		case "car is not available":
			c.JSON(http.StatusBadRequest, gin.H{"error": "Car is not available"})
		case "double_booking":
			c.JSON(http.StatusConflict, gin.H{"error": "Car is already booked for the selected dates"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking: " + err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Booking created successfully",
		"booking_id":  bookingID,
		"total_price": totalPrice,
		"invoice_url": invoiceURL,
	})
}

func UpdateBookingStatus(c *gin.Context) {
	ownerID := c.GetInt("user_id")
	bookingIDParam := c.Param("id")
	
	// import strconv manually handled later if needed, assume minimal inline or generic
	var bookingID int
	fmt.Sscanf(bookingIDParam, "%d", &bookingID) 

	var input models.BookingStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.UpdateBookingStatusByOwner(ownerID, bookingID, input.Status)
	if err != nil {
		switch err.Error() {
		case "booking not found":
			c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		case "forbidden":
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't own the car for this booking"})
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking status updated to " + input.Status})
}
