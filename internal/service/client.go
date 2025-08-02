package abacatepayService

import (
	"log"

	"github.com/go-resty/resty/v2"
)

const baseURL = "https://api.abacatepay.com/v1"

type ChargeRequest struct {
	PayerName     string  `json:"payer_name"`
	PayerDocument string  `json:"payer_document"`
	Amount        float64 `json:"amount"`
	Metadata      string  `json:"metadata"`
	CallbackURL   string  `json:"callback_url"`
}

func CreatePixCharge(token string, payload ChargeRequest) ([]byte, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+token).
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(baseURL + "/charges/pix/static")

	if err != nil {
		log.Println("Erro ao criar cobrança Pix:", err)
		return nil, err
	}

	return resp.Body(), nil
}
