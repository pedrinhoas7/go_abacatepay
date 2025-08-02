package abacatepayclient

type PixChargeRequest struct {
	PixKey       string `json:"pix_key"`
	Description  string `json:"description"`
	CustomerName string `json:"customer_name"`
	Amount       int    `json:"amount"`
}

type PixChargeResponse struct {
	ChargeID   string `json:"charge_id"`
	QrCode     string `json:"qr_code"`
	QrCodeUrl  string `json:"qr_code_url"`
	Expiration int    `json:"expiration"`
}
