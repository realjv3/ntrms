package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchATR_HappyPath(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{
			"meta": {"symbol": "AAPL"},
			"values": [{"datetime": "2024-01-01", "atr": "0.2500"}],
			"status": "ok"
		}`))
	}))
	defer srv.Close()

	origBase := BaseURL
	BaseURL = srv.URL
	defer func() { BaseURL = origBase }()

	// unit under test
	atr, err := FetchATR("fake-key", "AAPL")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if atr != 0.25 {
		t.Errorf("ATR = %f, want 0.25", atr)
	}
}

func TestFetchATR_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "error", "message": "Invalid API key"}`))
	}))
	defer srv.Close()

	origBase := BaseURL
	BaseURL = srv.URL
	defer func() { BaseURL = origBase }()

	// unit under test
	_, err := FetchATR("bad-key", "AAPL")
	if err == nil {
		t.Fatal("expected error for bad API key")
	}
}

func TestFetchPrice_HappyPath(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"price": "3.4200"}`))
	}))
	defer srv.Close()

	origBase := BaseURL
	BaseURL = srv.URL
	defer func() { BaseURL = origBase }()

	// unit under test
	price, err := FetchPrice("fake-key", "TSLA")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if price != 3.42 {
		t.Errorf("Price = %f, want 3.42", price)
	}
}

func TestFetchPrice_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "error", "message": "symbol not found"}`))
	}))
	defer srv.Close()

	origBase := BaseURL
	BaseURL = srv.URL
	defer func() { BaseURL = origBase }()

	// unit under test
	_, err := FetchPrice("fake-key", "INVALID")
	if err == nil {
		t.Fatal("expected error for invalid symbol")
	}
}

func TestFetchATR_EmptyValues(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "ok", "values": []}`))
	}))
	defer srv.Close()

	origBase := BaseURL
	BaseURL = srv.URL
	defer func() { BaseURL = origBase }()

	// unit under test
	_, err := FetchATR("fake-key", "EMPTY")
	if err == nil {
		t.Fatal("expected error for empty values")
	}
}
