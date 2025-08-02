package abacatepayclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	BaseURL   string
	SecretKey string
	http      *http.Client
}

func NewClient(baseURL, secretKey string) *Client {
	return &Client{
		BaseURL:   baseURL,
		SecretKey: secretKey,
		http:      &http.Client{},
	}
}

func (c *Client) doRequest(method, path string, body interface{}) ([]byte, error) {
	var buf io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, c.BaseURL+path, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.SecretKey)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func (c *Client) CreatePixCode(req PixChargeRequest) (any, error) {
	respBytes, err := c.doRequest("POST", "v1/pixQrCode/create", req)
	if err != nil {
		return nil, err
	}

	var res any
	err = json.Unmarshal(respBytes, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type CustomerRequest struct {
	Name      string `json:"name"`
	Cellphone string `json:"cellphone"`
	Email     string `json:"email"`
	TaxId     string `json:"taxId"`
}

func (c *Client) CreateCustomer(req CustomerRequest) (any, error) {
	fmt.Println("Creating customer with request: ", req)
	respBytes, err := c.doRequest("POST", "/v1/customer/create", req)
	if err != nil {
		fmt.Println("ERrr: ", err)
		return nil, err
	}

	var res any
	err = json.Unmarshal(respBytes, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type BillingProduct struct {
	ExternalID  string `json:"externalId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"` // em centavos
}

type BillingRequest struct {
	Frequency     string           `json:"frequency"` // "ONE_TIME"
	Methods       []string         `json:"methods"`   // ["PIX"]
	Products      []BillingProduct `json:"products"`
	ReturnUrl     string           `json:"returnUrl"`
	CompletionUrl string           `json:"completionUrl"`
	CustomerID    string           `json:"customerId"`
	AllowCoupons  bool             `json:"allowCoupons"`
	Coupons       []string         `json:"coupons,omitempty"`
}

func (c *Client) CreateBilling(req BillingRequest) (any, error) {
	respBytes, err := c.doRequest("POST", "/v1/billing/create", req)
	if err != nil {
		return nil, err
	}

	var res any
	err = json.Unmarshal(respBytes, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
