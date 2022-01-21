package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Balance struct {
	Symbol           string                    `json:"symbol"`
	AnnualReports    []BalanceAnnualReports    `json:"annualReports"`
	QuarterlyReports []BalanceQuarterlyReports `json:"quarterlyReports"`
}

type BalanceAnnualReports struct {
	FiscalDateEnding                       string `json:"fiscalDateEnding"`
	ReportedCurrency                       string `json:"reportedCurrency"`
	TotalAssets                            int    `json:"totalAssets,string"`
	TotalCurrentAssets                     int    `json:"totalCurrentAssets,string"`
	CashAndCashEquivalentsAtCarryingValue  int    `json:"cashAndCashEquivalentsAtCarryingValue,string"`
	CashAndShortTermInvestments            int    `json:"cashAndShortTermInvestments,string"`
	Inventory                              int    `json:"inventory,string"`
	CurrentNetReceivables                  int    `json:"currentNetReceivables,string"`
	TotalNonCurrentAssets                  int    `json:"totalNonCurrentAssets,string"`
	PropertyPlantEquipment                 int    `json:"propertyPlantEquipment,string"`
	AccumulatedDepreciationAmortizationPPE int    `json:"accumulatedDepreciationAmortizationPPE,string"`
	IntangibleAssets                       int    `json:"intangibleAssets,string"`
	IntangibleAssetsExcludingGoodwill      int    `json:"intangibleAssetsExcludingGoodwill,string"`
	Goodwill                               int    `json:"goodwill,string"`
	Investments                            int    `json:"investments,string"`
	LongTermInvestments                    int    `json:"longTermInvestments,string"`
	ShortTermInvestments                   int    `json:"shortTermInvestments,string"`
	OtherCurrentAssets                     int    `json:"otherCurrentAssets,string"`
	OtherNonCurrrentAssets                 int    `json:"otherNonCurrrentAssets,string"`
	TotalLiabilities                       int    `json:"totalLiabilities,string"`
	TotalCurrentLiabilities                int    `json:"totalCurrentLiabilities,string"`
	CurrentAccountsPayable                 int    `json:"currentAccountsPayable,string"`
	DeferredRevenue                        int    `json:"deferredRevenue,string"`
	CurrentDebt                            int    `json:"currentDebt,string"`
	ShortTermDebt                          int    `json:"shortTermDebt,string"`
	TotalNonCurrentLiabilities             int    `json:"totalNonCurrentLiabilities,string"`
	CapitalLeaseObligations                int    `json:"capitalLeaseObligations,string"`
	LongTermDebt                           int    `json:"longTermDebt,string"`
	CurrentLongTermDebt                    int    `json:"currentLongTermDebt,string"`
	LongTermDebtNoncurrent                 int    `json:"longTermDebtNoncurrent,string"`
	ShortLongTermDebtTotal                 int    `json:"shortLongTermDebtTotal,string"`
	OtherCurrentLiabilities                int    `json:"otherCurrentLiabilities,string"`
	OtherNonCurrentLiabilities             int    `json:"otherNonCurrentLiabilities,string"`
	TotalShareholderEquity                 int    `json:"totalShareholderEquity,string"`
	TreasuryStock                          int    `json:"treasuryStock,string"`
	RetainedEarnings                       int    `json:"retainedEarnings,string"`
	CommonStock                            int    `json:"commonStock,string"`
	CommonStockSharesOutstanding           int    `json:"commonStockSharesOutstanding,string"`
}

type BalanceQuarterlyReports struct {
	FiscalDateEnding                       string `json:"fiscalDateEnding"`
	ReportedCurrency                       string `json:"reportedCurrency"`
	TotalAssets                            int    `json:"totalAssets,string"`
	TotalCurrentAssets                     int    `json:"totalCurrentAssets,string"`
	CashAndCashEquivalentsAtCarryingValue  int    `json:"cashAndCashEquivalentsAtCarryingValue,string"`
	CashAndShortTermInvestments            int    `json:"cashAndShortTermInvestments,string"`
	Inventory                              int    `json:"inventory,string"`
	CurrentNetReceivables                  int    `json:"currentNetReceivables,string"`
	TotalNonCurrentAssets                  int    `json:"totalNonCurrentAssets,string"`
	PropertyPlantEquipment                 int    `json:"propertyPlantEquipment,string"`
	AccumulatedDepreciationAmortizationPPE int    `json:"accumulatedDepreciationAmortizationPPE,string"`
	IntangibleAssets                       int    `json:"intangibleAssets,string"`
	IntangibleAssetsExcludingGoodwill      int    `json:"intangibleAssetsExcludingGoodwill,string"`
	Goodwill                               int    `json:"goodwill,string"`
	Investments                            string `json:"investments"`
	LongTermInvestments                    int    `json:"longTermInvestments,string"`
	ShortTermInvestments                   int    `json:"shortTermInvestments,string"`
	OtherCurrentAssets                     int    `json:"otherCurrentAssets,string"`
	OtherNonCurrrentAssets                 int    `json:"otherNonCurrrentAssets,string"`
	TotalLiabilities                       int    `json:"totalLiabilities,string"`
	TotalCurrentLiabilities                int    `json:"totalCurrentLiabilities,string"`
	CurrentAccountsPayable                 int    `json:"currentAccountsPayable,string"`
	DeferredRevenue                        int    `json:"deferredRevenue,string"`
	CurrentDebt                            int    `json:"currentDebt,string"`
	ShortTermDebt                          int    `json:"shortTermDebt,string"`
	TotalNonCurrentLiabilities             int    `json:"totalNonCurrentLiabilities,string"`
	CapitalLeaseObligations                int    `json:"capitalLeaseObligations,string"`
	LongTermDebt                           int    `json:"longTermDebt,string"`
	CurrentLongTermDebt                    int    `json:"currentLongTermDebt,string"`
	LongTermDebtNoncurrent                 int    `json:"longTermDebtNoncurrent,string"`
	ShortLongTermDebtTotal                 int    `json:"shortLongTermDebtTotal,string"`
	OtherCurrentLiabilities                int    `json:"otherCurrentLiabilities,string"`
	OtherNonCurrentLiabilities             int    `json:"otherNonCurrentLiabilities,string"`
	TotalShareholderEquity                 int    `json:"totalShareholderEquity,string"`
	TreasuryStock                          int    `json:"treasuryStock,string"`
	RetainedEarnings                       int    `json:"retainedEarnings,string"`
	CommonStock                            int    `json:"commonStock,string"`
	CommonStockSharesOutstanding           int    `json:"commonStockSharesOutstanding,string"`
}

func BalanceSheetRequest(symbol string) (Balance, error) {
	const ENDPOINT_URL string = "BALANCE_SHEET"

	baseUrl := base_url()
	values := baseUrl.Query()

	values.Add("function", ENDPOINT_URL)
	values.Add("symbol", symbol)

	baseUrl.RawQuery = values.Encode()

	response, err := http.Get(baseUrl.String())

	if err != nil {
		return Balance{}, errors.New("the HTTP request has failed with an error")
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		balance := Balance{}

		json.Unmarshal(data, &balance)

		return balance, nil
	}
}
