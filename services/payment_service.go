package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"rent-car-project/config"
	"rent-car-project/models"
	"rent-car-project/repositories"
)

func CreateInvoice(bookingID int, amount float64) (string, error) {
	err := repositories.CreatePayment(bookingID, amount)
	if err != nil {
		return "", errors.New("failed to create payment record")
	}

	secretKey := config.GetXenditSecretKey()
	if secretKey == "" {
		return "", errors.New("xendit secret key not configured")
	}

	payload := models.XenditInvoiceRequest{
		ExternalID:  strconv.Itoa(bookingID),
		Amount:      amount,
		Description: fmt.Sprintf("Booking Car ID %d", bookingID),
	}

	jsonPayload, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "https://api.xendit.co/v2/invoices", bytes.NewBuffer(jsonPayload))
	req.SetBasicAuth(secretKey, "")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode >= 400 {
		return "", errors.New("failed to generate invoice from payment gateway")
	}
	defer resp.Body.Close()

	var invoiceResp models.XenditInvoiceResponse
	if err := json.NewDecoder(resp.Body).Decode(&invoiceResp); err != nil {
		return "", errors.New("invalid response from payment gateway")
	}

	return invoiceResp.InvoiceURL, nil
}

func HandleWebhook(callbackToken string, payload models.XenditWebhookPayload) error {
	expectedToken := config.GetXenditCallbackToken()
	if expectedToken == "" || callbackToken != expectedToken {
		return errors.New("invalid callback token")
	}

	bookingID, err := strconv.Atoi(payload.ExternalID)
	if err != nil {
		return errors.New("invalid external_id")
	}

	paidAt, _ := time.Parse(time.RFC3339, payload.PaidAt)
	return repositories.UpdatePaymentAndBookingStatus(bookingID, payload.Status, payload.ID, payload.PaymentMethod, paidAt)
}
