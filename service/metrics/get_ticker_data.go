package metrics

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type GetTickerResponse struct {
	CurrentPrice float64 `json:"ask"`
}

func request(url string) []byte {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalf("Error fetching data from %s", url)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading body from %s", url)
	}
	return body
}

func GetTickerData(ticker string) GetTickerResponse {
	url := os.Getenv("URL_TICKER")
	if url == "" {
		log.Fatalf("URL not found")
	}
	body := request(url + "/" + ticker)
	response := GetTickerResponse{}
	json.Unmarshal(body, &response)
	return response
}
