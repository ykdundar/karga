package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

type KAMA struct {
	MetaData              MetaData                              `json:"Meta Data"`
	TechnicalAnalysisKAMA map[string]TechnicalAnalysisKAMAValue `json:"Technical Analysis: KAMA"`
}
type MetaDataKama struct {
	Symbol        string `json:"1: Symbol"`
	Indicator     string `json:"2: Indicator"`
	LastRefreshed string `json:"3: Last Refreshed"`
	Interval      string `json:"4: Interval"`
	TimePeriod    int    `json:"5: Time Period"`
	SeriesType    string `json:"6: Series Type"`
	TimeZone      string `json:"7: Time Zone"`
}
type TechnicalAnalysisKAMAValue struct {
	Kama string `json:"KAMA"`
}

func KAMARequest(symbol string, interval string, time_period int, series_type string) (KAMA, error) {

	const ENDPOINT_URL string = "KAMA"
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
		return KAMA{}, errors.New("the HTTP request is failed with an error")
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		kama := KAMA{}

		json.Unmarshal(data, &kama)

		return kama, nil
	}
}
