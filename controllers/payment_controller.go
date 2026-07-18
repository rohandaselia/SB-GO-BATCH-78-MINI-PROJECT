package controllers

import (
	"net/http"
	"rent-car-project/models"
	"rent-car-project/services"

	"github.com/gin-gonic/gin"
)

func XenditWebhook(c *gin.Context) {
	callbackToken := c.GetHeader("x-callback-token")
	if callbackToken == "" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Missing callback token"})
		return
	}

	var payload models.XenditWebhookPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	err := services.HandleWebhook(callbackToken, payload)
	if err != nil {
		if err.Error() == "invalid callback token" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: invalid callback token"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process webhook"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Webhook received successfully"})
}
