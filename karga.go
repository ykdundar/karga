package main

import (
	"fmt"
	"net/url"
	"os"
)

var apikey string = os.Getenv("KARGA_TOKEN")

func base_url() *url.URL {
	return &url.URL{
		Scheme:     "https",
		Host:       "alphavantage.co",
		Path:       "query",
		ForceQuery: false,
		RawQuery:   fmt.Sprintf("apikey=%s", apikey),
	}
}

func overview(base_url *url.URL, symbol string) {
	const ENDPOINT_URL string = "OVERVIEW"

	values := base_url.Query()

	values.Add("function", ENDPOINT_URL)
	values.Add("symbol", symbol)

	base_url.RawQuery = values.Encode()

	fmt.Println(base_url)
}

func main() {
	overview(base_url(), "AAPL")
}
