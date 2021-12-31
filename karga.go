package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Overview struct {
	Symbol                      string `json:"Symbol"`
	AssetType                   string `json:"AssetType"`
	Name                        string `json:"Name"`
	Description                 string `json:"Description"`
	Cik                         string `json:"CIK"`
	Exchange                    string `json:"Exchange"`
	Currency                    string `json:"Currency"`
	Country                     string `json:"Country"`
	Sector                      string `json:"Sector"`
	Industry                    string `json:"Industry"`
	Address                     string `json:"Address"`
	FiscalYearEnd               string `json:"FiscalYearEnd"`
	LatestQuarter               string `json:"LatestQuarter"`
	MarketCapitalization        string `json:"MarketCapitalization"`
	Ebitda                      string `json:"EBITDA"`
	PERatio                     string `json:"PERatio"`
	PEGRatio                    string `json:"PEGRatio"`
	BookValue                   string `json:"BookValue"`
	DividendPerShare            string `json:"DividendPerShare"`
	DividendYield               string `json:"DividendYield"`
	Eps                         string `json:"EPS"`
	RevenuePerShareTTM          string `json:"RevenuePerShareTTM"`
	ProfitMargin                string `json:"ProfitMargin"`
	OperatingMarginTTM          string `json:"OperatingMarginTTM"`
	ReturnOnAssetsTTM           string `json:"ReturnOnAssetsTTM"`
	ReturnOnEquityTTM           string `json:"ReturnOnEquityTTM"`
	RevenueTTM                  string `json:"RevenueTTM"`
	GrossProfitTTM              string `json:"GrossProfitTTM"`
	DilutedEPSTTM               string `json:"DilutedEPSTTM"`
	QuarterlyEarningsGrowthYOY  string `json:"QuarterlyEarningsGrowthYOY"`
	QuarterlyRevenueGrowthYOY   string `json:"QuarterlyRevenueGrowthYOY"`
	AnalystTargetPrice          string `json:"AnalystTargetPrice"`
	TrailingPE                  string `json:"TrailingPE"`
	ForwardPE                   string `json:"ForwardPE"`
	PriceToSalesRatioTTM        string `json:"PriceToSalesRatioTTM"`
	PriceToBookRatio            string `json:"PriceToBookRatio"`
	EVToRevenue                 string `json:"EVToRevenue"`
	EVToEBITDA                  string `json:"EVToEBITDA"`
	Beta                        string `json:"Beta"`
	FityTwoWeekHigh             string `json:"52WeekHigh"`
	FiftyTwoWeekLow             string `json:"52WeekLow"`
	FifyDayMovingAverage        string `json:"50DayMovingAverage"`
	TwoHunderedDayMovingAverage string `json:"200DayMovingAverage"`
	SharesOutstanding           string `json:"SharesOutstanding"`
	DividendDate                string `json:"DividendDate"`
	ExDividendDate              string `json:"ExDividendDate"`
}

func base_url() *url.URL {
	return &url.URL{
		Scheme:     "https",
		Host:       "alphavantage.co",
		Path:       "query",
		ForceQuery: false,
		RawQuery:   fmt.Sprintf("apikey=%s", os.Getenv("KARGA_TOKEN")),
	}
}

func overview(symbol string) (Overview, error) {
	const ENDPOINT_URL string = "OVERVIEW"

	baseUrl := base_url()
	values := baseUrl.Query()

	values.Add("function", ENDPOINT_URL)
	values.Add("symbol", symbol)

	baseUrl.RawQuery = values.Encode()

	response, err := http.Get(baseUrl.String())

	if err != nil {
		return Overview{}, errors.New("the HTTP request is failed with an error")
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		overview := Overview{}

		json.Unmarshal(data, &overview)

		return overview, nil
	}
}

func main() {
	response, err := overview("IBM")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}
