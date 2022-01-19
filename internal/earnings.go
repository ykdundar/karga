package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Earnings struct {
	Symbol            string              `json:"symbol"`
	AnnualEarnings    []AnnualEarnings    `json:"annualEarnings"`
	QuarterlyEarnings []QuarterlyEarnings `json:"quarterlyEarnings"`
}
type AnnualEarnings struct {
	FiscalDateEnding string  `json:"fiscalDateEnding"`
	ReportedEPS      float64 `json:"reportedEPS,string"`
}
type QuarterlyEarnings struct {
	FiscalDateEnding   string  `json:"fiscalDateEnding"`
	ReportedDate       string  `json:"reportedDate"`
	ReportedEPS        float64 `json:"reportedEPS,string"`
	EstimatedEPS       float64 `json:"estimatedEPS,string"`
	Surprise           float64 `json:"surprise,string"`
	SurprisePercentage float64 `json:"surprisePercentage,string"`
}

func EarningsRequest(symbol string) (Earnings, error) {

	const ENDPOINT_URL string = "EARNINGS"
	baseUrl := base_url()
	values := baseUrl.Query()

	values.Add("function", ENDPOINT_URL)
	values.Add("symbol", symbol)

	baseUrl.RawQuery = values.Encode()

	response, err := http.Get(baseUrl.String())

	if err != nil {
		return Earnings{}, errors.New("the HTTP request is failed with an error")
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		earnings := Earnings{}

		json.Unmarshal(data, &earnings)

		return earnings, nil
	}
}
