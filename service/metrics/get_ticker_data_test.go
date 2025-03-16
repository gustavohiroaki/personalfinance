package metrics

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTickerData(t *testing.T) {

	ticker := "PETR4.SA"

	mockResponse := `{"currentPrice": 100.0}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("teste")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	os.Setenv("URL_TICKER", server.URL)
	defer server.Close()

	response := GetTickerData(ticker)
	assert.Equal(t, 100.0, response.CurrentPrice)
}
