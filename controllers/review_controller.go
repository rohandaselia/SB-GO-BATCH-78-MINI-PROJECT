package controllers

import (
	"fmt"
	"net/http"
	"rent-car-project/models"
	"rent-car-project/services"

	"github.com/gin-gonic/gin"
)

func CreateReview(c *gin.Context) {
	customerID := c.GetInt("user_id")
	bookingIDParam := c.Param("id")

	var bookingID int
	fmt.Sscanf(bookingIDParam, "%d", &bookingID)

	var input models.ReviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reviewID, err := services.AddReview(customerID, bookingID, input.Rating, input.Comment)
	if err != nil {
		switch err.Error() {
		case "booking not found":
			c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		case "forbidden":
			c.JSON(http.StatusForbidden, gin.H{"error": "You did not make this booking"})
		case "review already exists for this booking":
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":   "Review submitted successfully",
		"review_id": reviewID,
	})
}
