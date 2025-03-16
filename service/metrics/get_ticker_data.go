package metrics

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Response struct {
	CurrentPrice float64 `json:"ask"`
}

func GetTickerData(ticker string) Response {
	url := os.Getenv("URL_TICKER")
	if url == "" {
		log.Fatalf("URL not found")
	}

	resp, err := http.Get(url + "/" + ticker)

	if err != nil {
		log.Fatalf("Error fetching data from %s", url)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	response := Response{}
	json.Unmarshal(body, &response)
	return response
}
