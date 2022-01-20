package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ADX struct {
	MetaData             MetaDataADX                          `json:"Meta Data"`
	TechnicalAnalysisADX map[string]TechnicalAnalysisADXValue `json:"Technical Analysis: ADX"`
}

type MetaDataADX struct {
	Symbol        string `json:"1: Symbol"`
	Indicator     string `json:"2: Indicator"`
	LastRefreshed string `json:"3: Last Refreshed"`
	Interval      string `json:"4: Interval"`
	TimePeriod    int    `json:"5: Time Period"`
	TimeZone      string `json:"6: Time Zone"`
}

type TechnicalAnalysisADXValue struct {
	Adx string `json:"ADX"`
}

func ADXRequest(symbol string, interval string, time_period int) (ADX, error) {
	const ENDPOINT_URL string = "ADX"

	baseUrl := base_url()
	values := baseUrl.Query()

	values.Add("function", ENDPOINT_URL)
	values.Add("symbol", symbol)
	values.Add("interval", interval)
	values.Add("time_period", strconv.Itoa(time_period))

	baseUrl.RawQuery = values.Encode()

	response, err := http.Get(baseUrl.String())

	if err != nil {
		return ADX{}, errors.New("the HTTP request has failed with an error")
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		adx := ADX{}

		json.Unmarshal(data, &adx)

		return adx, nil
	}

}
