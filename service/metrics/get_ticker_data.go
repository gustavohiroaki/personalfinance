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

func GetTickerData(ticker string, transactionType string) Response {
	url := os.Getenv("URL_TICKER")
	if url == "" {
		log.Fatalf("URL not found")
	}
	switch transactionType {
	case "AÇÃO":
		ticker = ticker + ".SA"
	case "FII":
		ticker = ticker + ".SA"
	case "ETF":
		ticker = ticker + ".SA"
	default:
	}
	resp, err := http.Get(url + "/" + ticker)

	if err != nil {
		log.Fatalf("Error fetching data from %s", url)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading body from %s", url)
	}
	response := Response{}
	json.Unmarshal(body, &response)
	println(response.CurrentPrice)
	return response
}
