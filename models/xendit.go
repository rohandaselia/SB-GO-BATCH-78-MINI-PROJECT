package models

type XenditInvoiceRequest struct {
	ExternalID  string  `json:"external_id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

type XenditInvoiceResponse struct {
	ID         string `json:"id"`
	ExternalID string `json:"external_id"`
	Status     string `json:"status"`
	InvoiceURL string `json:"invoice_url"`
}

type XenditWebhookPayload struct {
	ID            string `json:"id"`
	ExternalID    string `json:"external_id"`
	Status        string `json:"status"`
	PaidAmount    float64 `json:"paid_amount"`
	PaymentMethod string `json:"payment_method"`
	PaidAt        string `json:"paid_at"`
}
