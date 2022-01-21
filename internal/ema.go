package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

type EMA struct {
	MetaData             MetaDataEma                           `json:"Meta Data"`
	TechnicalAnalysisEMA map[string]TechnicalAnalysisEMAValues `json:"Technical Analysis: EMA"`
}

type MetaDataEma struct {
	Symbol        string `json:"1: Symbol"`
	Indicator     string `json:"2: Indicator"`
	LastRefreshed string `json:"3: Last Refreshed"`
	Interval      string `json:"4: Interval"`
	TimePeriod    int    `json:"5: Time Period"`
	SeriesType    string `json:"6: Series Type"`
	TimeZone      string `json:"7: Time Zone"`
}

type TechnicalAnalysisEMAValues struct {
	Ema string `json:"EMA"`
}

func EMARequest(symbol string, interval string, timePeriod int, series_type string) (EMA, error) {
	const ENDPOINT_URL string = "EMA"

	baseUrl := base_url()
	values := baseUrl.Query()

	values.Add("function", ENDPOINT_URL)
	values.Add("symbol", symbol)
	values.Add("interval", interval)
	values.Add("time_period", strconv.Itoa(timePeriod))
	values.Add("series_type", series_type)

	baseUrl.RawQuery = values.Encode()

	response, err := http.Get(baseUrl.String())

	if err != nil {
		return EMA{}, errors.New("the HTTP request has failed with an error")
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		ema := EMA{}

		json.Unmarshal(data, &ema)

		return ema, nil
	}
}
