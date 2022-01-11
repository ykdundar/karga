package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
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
func rsi(symbol string, interval string, time_period int, series_type string) (RSIResponse, error) {
	const ENDPOINT_URL string = "RSI"
	baseUrl := base_url()
	values := baseUrl.Query()

	values.Add("function", ENDPOINT_URL)
	values.Add("symbol", symbol)
	values.Add("interval", interval)
	values.Add("time_period", strconv.Itoa(time_period))
	values.Add("series_type", series_type)

	baseUrl.RawQuery = values.Encode()

	response, err := http.Get(baseUrl.String())

	if err != nil {
		return RSIResponse{}, errors.New("the HTTP request is failed with an error")
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		rsi := RSIResponse{}

		json.Unmarshal(data, &rsi)

		return rsi, nil
	}

}

func globalQuote(symbol string) (GlobalQuote, error) {
	const ENDPOINT_URL string = "GLOBAL_QUOTE"

	baseUrl := base_url()
	values := baseUrl.Query()

	values.Add("function", ENDPOINT_URL)
	values.Add("symbol", symbol)

	baseUrl.RawQuery = values.Encode()

	response, err := http.Get(baseUrl.String())

	if err != nil {
		return GlobalQuote{}, errors.New("the HTTP request is failed with an error")
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		globalQuote := GlobalQuote{}

		json.Unmarshal(data, &globalQuote)

		return globalQuote, nil
	}

}

func overview(symbol string) (Overview, error) {
	const ENDPOINT_URL string = "OVERVIEW"

	baseUrl := base_url()
	values := baseUrl.Query()

	values.Add("function", ENDPOINT_URL)
	values.Add("symbol", symbol)

	baseUrl.RawQuery = values.Encode()

	response, err := http.Get(baseUrl.String())

	if err != nil {
		return Overview{}, errors.New("the HTTP request is failed with an error")
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		overview := Overview{}

		json.Unmarshal(data, &overview)

		return overview, nil
	}
}

func main() {
	response, err := rsi("IBM", "weekly", 10, "open")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}
