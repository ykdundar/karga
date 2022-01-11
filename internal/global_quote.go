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
	Symbol           string `json:"01. symbol"`
	Open             string `json:"02. open"`
	High             string `json:"03. high"`
	Low              string `json:"04. low"`
	Price            string `json:"05. price"`
	Volume           string `json:"06. volume"`
	LatestTradingDay string `json:"07. latest trading day"`
	PreviousClose    string `json:"08. previous close"`
	Change           string `json:"09. change"`
	ChangePercent    string `json:"10. change percent"`
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
