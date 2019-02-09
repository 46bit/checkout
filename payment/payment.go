package payment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Request
// POST https://publicapi.payments.service.gov.uk/v1/payments
// Content-Type: application/json
// Accept: application/json
// Authorization: $API_KEY
// {
//   "amount": $number-of-pence,
//   "reference": "$random-value-per-payment",
//   "return_url": "https://github.com/46bit/checkout",
//   "description": "Checkout payment",
//   "language": "en"
// }
//
// Response:
// - Verify 200 status code
// - Return `_links.next_url.href` from JSON payload

func PaymentLink(price uint, id, apiKey string) (string, error) {
	apiUrl := "https://publicapi.payments.service.gov.uk/v1/payments"
	bodyMap := map[string]interface{}{
		"amount":      price,
		"reference":   id,
		"return_url":  "https://github.com/46bit/checkout",
		"description": "Checkout payment",
		"language":    "en",
	}
	body, err := json.Marshal(bodyMap)
	if err != nil {
		return "", err
	}
	request, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	if response.StatusCode != 201 {
		return "", fmt.Errorf("non-201 status code '%d'", response.StatusCode)
	}
	respBodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	var responseBody struct {
		Links struct {
			NextUrl struct {
				Href string `json:"href"`
			} `json:"next_url"`
		} `json:"_links"`
	}
	err = json.Unmarshal(respBodyBytes, &responseBody)
	if err != nil {
		return "", err
	}
	return responseBody.Links.NextUrl.Href, nil
}
