package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type GlobalQuote struct {
	GlobalQuoteDetails GlobalQuoteDetails `json:"Global Quote"`
}

type GlobalQuoteDetails struct {
	Symbol           string  `json:"01. symbol"`
	Open             float64 `json:"02. open,string"`
	High             float64 `json:"03. high,string"`
	Low              float64 `json:"04. low,string"`
	Price            float64 `json:"05. price,string"`
	Volume           int     `json:"06. volume,string"`
	LatestTradingDay string  `json:"07. latest trading day"`
	PreviousClose    float64 `json:"08. previous close,string"`
	Change           float64 `json:"09. change,string"`
	ChangePercent    float64 `json:"10. change percent,string"`
}

func GlobalQuoteRequest(symbol string) (GlobalQuote, error) {
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
