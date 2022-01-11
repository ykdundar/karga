package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type TimeSeriesValue struct {
	OneOpen    string `json:"1. open, omitempty"`
	TwoHigh    string `json:"2. high, omitempty"`
	ThreeLow   string `json:"3. low", omitempty`
	FourClose  string `json:"4. close, omitempty"`
	FiveVolume string `json:"5. volume, omitempty"`
}

type Intraday struct {
	MetaData struct {
		OneInformation     string `json:"1. Information"`
		TwoSymbol          string `json:"2. Symbol"`
		ThreeLastRefreshed string `json:"3. Last Refreshed"`
		FourInterval       string `json:"4. Interval"`
		FiveOutputSize     string `json:"5. Output Size"`
		SixTimeZone        string `json:"6. Time Zone"`
	} `json:"Meta Data"`
	TimeSeries1Min  map[string]TimeSeriesValue `json:"Time Series (1min)"`
	TimeSeries5Min  map[string]TimeSeriesValue `json:"Time Series (5min)"`
	TimeSeries15Min map[string]TimeSeriesValue `json:"Time Series (15min)"`
	TimeSeries30Min map[string]TimeSeriesValue `json:"Time Series (30min)"`
	TimeSeries60Min map[string]TimeSeriesValue `json:"Time Series (60min)"`
}

func IntradayRequest(symbol, interval string) (Intraday, error) {
	const ENDPOINT_URL string = "TIME_SERIES_INTRADAY"
	baseUrl := base_url()
	values := baseUrl.Query()

	values.Add("function", ENDPOINT_URL)
	values.Add("symbol", symbol)
	values.Add("interval", interval)

	baseUrl.RawQuery = values.Encode()

	response, err := http.Get(baseUrl.String())

	if err != nil {
		return Intraday{}, errors.New("the HTTP request is failed with an error")
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		intraday := Intraday{}

		json.Unmarshal(data, &intraday)

		return intraday, nil
	}
}
