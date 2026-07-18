package controllers

import (
	"rent-car-project/models"
	"net/http"
	"rent-car-project/services"
	"strconv"

	"github.com/gin-gonic/gin"
)


func CreateCar(c *gin.Context) {
	ownerID := c.GetInt("user_id")
	var input models.CarInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	carID, err := services.CreateCar(ownerID, input.Brand, input.Model, input.LicensePlate, input.Description, input.PricePerDay)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create car: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Car added successfully", "car_id": carID})
}

func GetMyCars(c *gin.Context) {
	ownerID := c.GetInt("user_id")
	cars, err := services.GetMyCars(ownerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cars"})
		return
	}
	if cars == nil {
		cars = []models.Car{} // import models not needed if we handle slice here or in service properly, but let's just return empty array
	}
	c.JSON(http.StatusOK, cars)
}

func UpdateCar(c *gin.Context) {
	ownerID := c.GetInt("user_id")
	carID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
		return
	}

	var input models.CarInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.UpdateCar(ownerID, carID, input.Brand, input.Model, input.LicensePlate, input.Description, input.PricePerDay)
	if err != nil {
		switch err.Error() {
		case "car not found":
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case "forbidden":
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't own this car"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car updated successfully"})
}

func DeleteCar(c *gin.Context) {
	ownerID := c.GetInt("user_id")
	carID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
		return
	}

	err = services.DeleteCar(ownerID, carID)
	if err != nil {
		if err.Error() == "not found or forbidden" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Car not found or you don't have permission"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete car"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully"})
}
