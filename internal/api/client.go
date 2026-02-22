package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// BaseURL is the Twelve Data API root. Override in tests.
var BaseURL = "https://api.twelvedata.com"

var httpClient = &http.Client{Timeout: 10 * time.Second}

// atrResponse models the Twelve Data /atr endpoint JSON.
type atrResponse struct {
	Status string `json:"status"`
	Values []struct {
		ATR string `json:"atr"`
	} `json:"values"`
	Message string `json:"message"`
}

// priceResponse models the Twelve Data /price endpoint JSON.
type priceResponse struct {
	Price   string `json:"price"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// FetchATR retrieves the most recent 14-day daily ATR for the given ticker.
func FetchATR(apiKey, ticker string) (float64, error) {
	url := fmt.Sprintf("%s/atr?symbol=%s&interval=1day&time_period=14&apikey=%s",
		BaseURL, ticker, apiKey)

	resp, err := httpClient.Get(url)
	if err != nil {
		return 0, fmt.Errorf("atr request failed: %w", err)
	}
	defer resp.Body.Close()

	var data atrResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, fmt.Errorf("atr decode failed: %w", err)
	}

	if data.Status == "error" {
		return 0, fmt.Errorf("atr api error: %s", data.Message)
	}

	if len(data.Values) == 0 {
		return 0, fmt.Errorf("atr response contained no values")
	}

	atr, err := strconv.ParseFloat(data.Values[0].ATR, 64)
	if err != nil {
		return 0, fmt.Errorf("atr parse failed: %w", err)
	}

	return atr, nil
}

// FetchPrice retrieves the current price for the given ticker.
func FetchPrice(apiKey, ticker string) (float64, error) {
	url := fmt.Sprintf("%s/price?symbol=%s&apikey=%s",
		BaseURL, ticker, apiKey)

	resp, err := httpClient.Get(url)
	if err != nil {
		return 0, fmt.Errorf("price request failed: %w", err)
	}
	defer resp.Body.Close()

	var data priceResponse

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, fmt.Errorf("price decode failed: %w", err)
	}

	if data.Status == "error" {
		return 0, fmt.Errorf("price api error: %s", data.Message)
	}

	price, err := strconv.ParseFloat(data.Price, 64)
	if err != nil {
		return 0, fmt.Errorf("price parse failed: %w", err)
	}

	return price, nil
}
