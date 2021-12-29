package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func base_url() *url.URL {
	return &url.URL{
		Scheme:     "https",
		Host:       "alphavantage.co",
		Path:       "query",
		ForceQuery: false,
		RawQuery:   fmt.Sprintf("apikey=%s", os.Getenv("KARGA_TOKEN")),
	}
}

func overview(symbol string) string {
	const ENDPOINT_URL string = "OVERVIEW"

	baseUrl := base_url()
	values := baseUrl.Query()

	values.Add("function", ENDPOINT_URL)
	values.Add("symbol", symbol)

	baseUrl.RawQuery = values.Encode()

	response, err := http.Get(baseUrl.String())
	if err != nil {
		return fmt.Sprintf("The HTTP request is failed with an error %s\n ", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		return string(data)
	}
}

func main() {
	fmt.Println(overview("IBM"))
}
