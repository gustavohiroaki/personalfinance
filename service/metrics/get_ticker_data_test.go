package metrics

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBrazilianTickerData(t *testing.T) {
	ticker := "PETR4"
	transactionType := "AÇÃO"

	mockResponse := `{"ask": 100.0}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.URL.Path, "/PETR4.SA")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	os.Setenv("URL_TICKER", server.URL)
	defer server.Close()

	response := GetTickerData(ticker, transactionType)
	assert.Equal(t, 100.0, response.CurrentPrice)
}

func TestGetNonBrazilianTickerData(t *testing.T) {
	ticker := "DIS"
	transactionType := "STOCK"

	mockResponse := `{"ask": 100.0}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.URL.Path, "/DIS")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	os.Setenv("URL_TICKER", server.URL)
	defer server.Close()

	response := GetTickerData(ticker, transactionType)
	assert.Equal(t, 100.0, response.CurrentPrice)
}
