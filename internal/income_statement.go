package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type IncomeStatement struct {
	Symbol           string             `json:"symbol"`
	AnnualReports    []AnnualReports    `json:"annualReports"`
	QuarterlyReports []QuarterlyReports `json:"quarterlyReports"`
}
type AnnualReports struct {
	FiscalDateEnding                  string `json:"fiscalDateEnding,string"`
	ReportedCurrency                  string `json:"reportedCurrency,string"`
	GrossProfit                       int    `json:"grossProfit,string"`
	TotalRevenue                      int    `json:"totalRevenue,string"`
	CostOfRevenue                     int    `json:"costOfRevenue,string"`
	CostofGoodsAndServicesSold        int    `json:"costofGoodsAndServicesSold,string"`
	OperatingIncome                   int    `json:"operatingIncome,string"`
	SellingGeneralAndAdministrative   int    `json:"sellingGeneralAndAdministrative,string"`
	ResearchAndDevelopment            int    `json:"researchAndDevelopment,string"`
	OperatingExpenses                 int    `json:"operatingExpenses,string"`
	InvestmentIncomeNet               string `json:"investmentIncomeNet"`
	NetInterestIncome                 int    `json:"netInterestIncome,string"`
	InterestIncome                    int    `json:"interestIncome,string"`
	InterestExpense                   int    `json:"interestExpense,string"`
	NonInterestIncome                 string `json:"nonInterestIncome"`
	OtherNonOperatingIncome           int    `json:"otherNonOperatingIncome,string"`
	Depreciation                      int    `json:"depreciation,string"`
	DepreciationAndAmortization       int    `json:"depreciationAndAmortization,string"`
	IncomeBeforeTax                   int    `json:"incomeBeforeTax,string"`
	IncomeTaxExpense                  int    `json:"incomeTaxExpense,string"`
	InterestAndDebtExpense            int    `json:"interestAndDebtExpense,string"`
	NetIncomeFromContinuingOperations int    `json:"netIncomeFromContinuingOperations,string"`
	ComprehensiveIncomeNetOfTax       int    `json:"comprehensiveIncomeNetOfTax,string"`
	Ebit                              int    `json:"ebit,string"`
	Ebitda                            int    `json:"ebitda,string"`
	NetIncome                         int    `json:"netIncome,string"`
}
type QuarterlyReports struct {
	FiscalDateEnding                  string `json:"fiscalDateEnding,string"`
	ReportedCurrency                  string `json:"reportedCurrency,string"`
	GrossProfit                       int    `json:"grossProfit,string"`
	TotalRevenue                      int    `json:"totalRevenue,string"`
	CostOfRevenue                     int    `json:"costOfRevenue,string"`
	CostofGoodsAndServicesSold        int    `json:"costofGoodsAndServicesSold,string"`
	OperatingIncome                   int    `json:"operatingIncome,string"`
	SellingGeneralAndAdministrative   int    `json:"sellingGeneralAndAdministrative,string"`
	ResearchAndDevelopment            int    `json:"researchAndDevelopment,string"`
	OperatingExpenses                 int    `json:"operatingExpenses,string"`
	InvestmentIncomeNet               string `json:"investmentIncomeNet"`
	NetInterestIncome                 int    `json:"netInterestIncome,string"`
	InterestIncome                    int    `json:"interestIncome,string"`
	InterestExpense                   int    `json:"interestExpense,string"`
	NonInterestIncome                 string `json:"nonInterestIncome"`
	OtherNonOperatingIncome           int    `json:"otherNonOperatingIncome,string"`
	Depreciation                      int    `json:"depreciation,string"`
	DepreciationAndAmortization       int    `json:"depreciationAndAmortization,string"`
	IncomeBeforeTax                   int    `json:"incomeBeforeTax,string"`
	IncomeTaxExpense                  int    `json:"incomeTaxExpense,string"`
	InterestAndDebtExpense            int    `json:"interestAndDebtExpense,string"`
	NetIncomeFromContinuingOperations int    `json:"netIncomeFromContinuingOperations,string"`
	ComprehensiveIncomeNetOfTax       int    `json:"comprehensiveIncomeNetOfTax,string"`
	Ebit                              int    `json:"ebit,string"`
	Ebitda                            int    `json:"ebitda,string"`
	NetIncome                         int    `json:"netIncome,string"`
}

func IncomeStatementRequest(symbol string) (IncomeStatement, error) {

	const ENDPOINT_URL string = "INCOME_STATEMENT"

	baseUrl := base_url()
	values := baseUrl.Query()

	values.Add("function", ENDPOINT_URL)
	values.Add("symbol", symbol)

	baseUrl.RawQuery = values.Encode()

	response, err := http.Get(baseUrl.String())

	if err != nil {
		return IncomeStatement{}, errors.New("the HTTP request is failed with an error")
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		incomeStatement := IncomeStatement{}

		json.Unmarshal(data, &incomeStatement)

		return incomeStatement, nil
	}
}
