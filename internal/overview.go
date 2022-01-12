package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Overview struct {
	Symbol                      string  `json:"Symbol"`
	AssetType                   string  `json:"AssetType"`
	Name                        string  `json:"Name"`
	Description                 string  `json:"Description"`
	Cik                         int     `json:"CIK,string"`
	Exchange                    string  `json:"Exchange"`
	Currency                    string  `json:"Currency"`
	Country                     string  `json:"Country"`
	Sector                      string  `json:"Sector"`
	Industry                    string  `json:"Industry"`
	Address                     string  `json:"Address"`
	FiscalYearEnd               string  `json:"FiscalYearEnd"`
	LatestQuarter               string  `json:"LatestQuarter"`
	MarketCapitalization        int     `json:"MarketCapitalization,string"`
	Ebitda                      int     `json:"EBITDA,string"`
	PERatio                     float64 `json:"PERatio,string"`
	PEGRatio                    float64 `json:"PEGRatio,string"`
	BookValue                   float64 `json:"BookValue,string"`
	DividendPerShare            float64 `json:"DividendPerShare,string"`
	DividendYield               float64 `json:"DividendYield,string"`
	Eps                         float64 `json:"EPS,string"`
	RevenuePerShareTTM          float64 `json:"RevenuePerShareTTM,string"`
	ProfitMargin                float64 `json:"ProfitMargin,string"`
	OperatingMarginTTM          float64 `json:"OperatingMarginTTM,string"`
	ReturnOnAssetsTTM           float64 `json:"ReturnOnAssetsTTM,string"`
	ReturnOnEquityTTM           float64 `json:"ReturnOnEquityTTM,string"`
	RevenueTTM                  int     `json:"RevenueTTM,string"`
	GrossProfitTTM              int     `json:"GrossProfitTTM,string"`
	DilutedEPSTTM               float64 `json:"DilutedEPSTTM,string"`
	QuarterlyEarningsGrowthYOY  float64 `json:"QuarterlyEarningsGrowthYOY,string"`
	QuarterlyRevenueGrowthYOY   float64 `json:"QuarterlyRevenueGrowthYOY,string"`
	AnalystTargetPrice          float64 `json:"AnalystTargetPrice,string"`
	TrailingPE                  float64 `json:"TrailingPE,string"`
	ForwardPE                   float64 `json:"ForwardPE,string"`
	PriceToSalesRatioTTM        float64 `json:"PriceToSalesRatioTTM,string"`
	PriceToBookRatio            float64 `json:"PriceToBookRatio,string"`
	EVToRevenue                 float64 `json:"EVToRevenue,string"`
	EVToEBITDA                  float64 `json:"EVToEBITDA,string"`
	Beta                        float64 `json:"Beta,string"`
	FityTwoWeekHigh             float64 `json:"52WeekHigh,string"`
	FiftyTwoWeekLow             float64 `json:"52WeekLow,string"`
	FifyDayMovingAverage        float64 `json:"50DayMovingAverage,string"`
	TwoHunderedDayMovingAverage float64 `json:"200DayMovingAverage,string"`
	SharesOutstanding           int     `json:"SharesOutstanding,string"`
	DividendDate                string  `json:"DividendDate"`
	ExDividendDate              string  `json:"ExDividendDate"`
}

func OverviewRequest(symbol string) (Overview, error) {
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
