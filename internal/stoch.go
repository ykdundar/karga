package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Stoch struct {
	MetaData               MetaDataStoch                          `json:"Meta Data"`
	TechnicalAnalysisSTOCH map[string]TechnicalAnalysisSTOCHValue `json:"Technical Analysis: STOCH"`
}

type MetaDataStoch struct {
	Symbol        string `json:"1: Symbol"`
	Indicator     string `json:"2: Indicator"`
	LastRefreshed string `json:"3: Last Refreshed"`
	Interval      string `json:"4: Interval"`
	FastKPeriod   int    `json:"5.1: FastK Period"`
	SlowKPeriod   int    `json:"5.2: SlowK Period"`
	SlowKMAType   int    `json:"5.3: SlowK MA Type"`
	SlowDPeriod   int    `json:"5.4: SlowD Period"`
	SlowDMAType   int    `json:"5.5: SlowD MA Type"`
	TimeZone      string `json:"6: Time Zone"`
}

type TechnicalAnalysisSTOCHValue struct {
	SlowK string `json:"SlowK"`
}

func StochRequest(symbol, interval string) (Stoch, error) {
	const ENDPOINT_URL string = "STOCH"

	baseUrl := base_url()
	values := baseUrl.Query()

	values.Add("function", ENDPOINT_URL)
	values.Add("symbol", symbol)
	values.Add("interval", interval)

	baseUrl.RawQuery = values.Encode()

	response, err := http.Get(baseUrl.String())

	if err != nil {
		return Stoch{}, errors.New("the HTTP request has failed with an error")
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		stoch := Stoch{}

		json.Unmarshal(data, &stoch)

		return stoch, nil
	}

}
