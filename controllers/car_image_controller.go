package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"rent-car-project/services"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadCarImage(c *gin.Context) {
	ownerID := c.GetInt("user_id")
	carIDParam := c.Param("id")
	carID, err := strconv.Atoi(carIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only JPG and PNG are allowed"})
		return
	}

	isPrimaryStr := c.PostForm("is_primary")
	isPrimary := false
	if isPrimaryStr == "true" || isPrimaryStr == "1" {
		isPrimary = true
	}

	fileName := fmt.Sprintf("car-%d-%d%s", carID, time.Now().Unix(), ext)
	savePath := filepath.Join("public/images", fileName)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	host := c.Request.Host
	scheme := "http"
	if c.Request.TLS != nil || c.GetHeader("X-Forwarded-Proto") == "https" {
		scheme = "https"
	}
	imageURL := fmt.Sprintf("%s://%s/public/images/%s", scheme, host, fileName)

	imageID, err := services.AddCarImage(ownerID, carID, imageURL, isPrimary)
	if err != nil {
		switch err.Error() {
		case "car not found":
			c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		case "forbidden":
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't own this car"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":   "Image uploaded successfully",
		"image_id":  imageID,
		"image_url": imageURL,
	})
}
