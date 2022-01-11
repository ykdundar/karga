package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

type RSIValue struct {
	Rsi string `json:"RSI"`
}

type RSIResponse struct {
	MetaData struct {
		OneSymbol          string `json:"1: Symbol"`
		TwoIndicator       string `json:"2: Indicator"`
		ThreeLastRefreshed string `json:"3: Last Refreshed"`
		FourInterval       string `json:"4: Interval"`
		FiveTimePeriod     int    `json:"5: Time Period"`
		SixSeriesType      string `json:"6: Series Type"`
		SevenTimeZone      string `json:"7: Time Zone"`
	} `json:"Meta Data"`
	TechnicalAnalysisRSI map[string]RSIValue `json:"Technical Analysis: RSI"`
}

func RSIRequest(symbol string, interval string, time_period int, series_type string) (RSIResponse, error) {
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