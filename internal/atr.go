package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ATR struct {
	MetaData             MetaData                              `json:"Meta Data"`
	TechnicalAnalysisATR map[string]TechnicalAnalysisATRValues `json:"Technical Analysis: ATR"`
}

type MetaData struct {
	Symbol        string `json:"1: Symbol"`
	Indicator     string `json:"2: Indicator"`
	LastRefreshed string `json:"3: Last Refreshed"`
	Interval      string `json:"4: Interval"`
	TimePeriod    int    `json:"5: Time Period"`
	TimeZone      string `json:"6: Time Zone"`
}

type TechnicalAnalysisATRValues struct {
	Atr float64 `json:"ATR,string"`
}

func ATRRequest(symbol string, interval string, time_period int) (ATR, error) {
	const ENDPOINT_URL string = "ATR"

	baseUrl := base_url()
	values := baseUrl.Query()

	values.Add("function", ENDPOINT_URL)
	values.Add("symbol", symbol)
	values.Add("interval", interval)
	values.Add("time_period", strconv.Itoa(time_period))

	baseUrl.RawQuery = values.Encode()

	response, err := http.Get(baseUrl.String())

	if err != nil {
		return ATR{}, errors.New("the HTTP request is failed with an error")
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		atr := ATR{}

		json.Unmarshal(data, &atr)

		return atr, nil
	}

}
