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
	Zero1Symbol           string `json:"01. symbol"`
	Zero2Open             string `json:"02. open"`
	Zero3High             string `json:"03. high"`
	Zero4Low              string `json:"04. low"`
	Zero5Price            string `json:"05. price"`
	Zero6Volume           string `json:"06. volume"`
	Zero7LatestTradingDay string `json:"07. latest trading day"`
	Zero8PreviousClose    string `json:"08. previous close"`
	Zero9Change           string `json:"09. change"`
	One0ChangePercent     string `json:"10. change percent"`
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
