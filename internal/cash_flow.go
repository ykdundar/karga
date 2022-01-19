package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type CashFlow struct {
	Symbol           string                     `json:"symbol"`
	AnnualReports    []CashFlowAnnualReports    `json:"annualReports"`
	QuarterlyReports []CashFlowQuarterlyReports `json:"quarterlyReports"`
}
type CashFlowAnnualReports struct {
	FiscalDateEnding                                          string `json:"fiscalDateEnding"`
	ReportedCurrency                                          string `json:"reportedCurrency"`
	OperatingCashflow                                         int    `json:"operatingCashflow,string,string"`
	PaymentsForOperatingActivities                            int    `json:"paymentsForOperatingActivities,string"`
	ProceedsFromOperatingActivities                           string `json:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities                              int    `json:"changeInOperatingLiabilities,string"`
	ChangeInOperatingAssets                                   int    `json:"changeInOperatingAssets,string"`
	DepreciationDepletionAndAmortization                      int    `json:"depreciationDepletionAndAmortization,string"`
	CapitalExpenditures                                       int    `json:"capitalExpenditures,string"`
	ChangeInReceivables                                       int    `json:"changeInReceivables,string"`
	ChangeInInventory                                         int    `json:"changeInInventory,string"`
	ProfitLoss                                                int    `json:"profitLoss,string"`
	CashflowFromInvestment                                    int    `json:"cashflowFromInvestment,string"`
	CashflowFromFinancing                                     int    `json:"cashflowFromFinancing,string"`
	ProceedsFromRepaymentsOfShortTermDebt                     int    `json:"proceedsFromRepaymentsOfShortTermDebt,string"`
	PaymentsForRepurchaseOfCommonStock                        string `json:"paymentsForRepurchaseOfCommonStock"`
	PaymentsForRepurchaseOfEquity                             string `json:"paymentsForRepurchaseOfEquity"`
	PaymentsForRepurchaseOfPreferredStock                     string `json:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                                            int    `json:"dividendPayout,string"`
	DividendPayoutCommonStock                                 int    `json:"dividendPayoutCommonStock,string"`
	DividendPayoutPreferredStock                              string `json:"dividendPayoutPreferredStock"`
	ProceedsFromIssuanceOfCommonStock                         string `json:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet int    `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet,string"`
	ProceedsFromIssuanceOfPreferredStock                      string `json:"proceedsFromIssuanceOfPreferredStock"`
	ProceedsFromRepurchaseOfEquity                            int    `json:"proceedsFromRepurchaseOfEquity,string"`
	ProceedsFromSaleOfTreasuryStock                           string `json:"proceedsFromSaleOfTreasuryStock"`
	ChangeInCashAndCashEquivalents                            int    `json:"changeInCashAndCashEquivalents,string"`
	ChangeInExchangeRate                                      string `json:"changeInExchangeRate"`
	NetIncome                                                 int    `json:"netIncome,string"`
}
type CashFlowQuarterlyReports struct {
	FiscalDateEnding                                          string `json:"fiscalDateEnding"`
	ReportedCurrency                                          string `json:"reportedCurrency"`
	OperatingCashflow                                         int    `json:"operatingCashflow,string"`
	PaymentsForOperatingActivities                            string `json:"paymentsForOperatingActivities"`
	ProceedsFromOperatingActivities                           string `json:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities                              string `json:"changeInOperatingLiabilities"`
	ChangeInOperatingAssets                                   string `json:"changeInOperatingAssets"`
	DepreciationDepletionAndAmortization                      int    `json:"depreciationDepletionAndAmortization,string"`
	CapitalExpenditures                                       int    `json:"capitalExpenditures,string"`
	ChangeInReceivables                                       string `json:"changeInReceivables"`
	ChangeInInventory                                         string `json:"changeInInventory"`
	ProfitLoss                                                int    `json:"profitLoss,string"`
	CashflowFromInvestment                                    int    `json:"cashflowFromInvestment,string"`
	CashflowFromFinancing                                     int    `json:"cashflowFromFinancing,string"`
	ProceedsFromRepaymentsOfShortTermDebt                     int    `json:"proceedsFromRepaymentsOfShortTermDebt,string"`
	PaymentsForRepurchaseOfCommonStock                        int    `json:"paymentsForRepurchaseOfCommonStock,string"`
	PaymentsForRepurchaseOfEquity                             int    `json:"paymentsForRepurchaseOfEquity,string"`
	PaymentsForRepurchaseOfPreferredStock                     string `json:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                                            int    `json:"dividendPayout,string"`
	DividendPayoutCommonStock                                 int    `json:"dividendPayoutCommonStock,string"`
	DividendPayoutPreferredStock                              string `json:"dividendPayoutPreferredStock"`
	ProceedsFromIssuanceOfCommonStock                         string `json:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet int    `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet,string"`
	ProceedsFromIssuanceOfPreferredStock                      string `json:"proceedsFromIssuanceOfPreferredStock"`
	ProceedsFromRepurchaseOfEquity                            int    `json:"proceedsFromRepurchaseOfEquity,string"`
	ProceedsFromSaleOfTreasuryStock                           string `json:"proceedsFromSaleOfTreasuryStock"`
	ChangeInCashAndCashEquivalents                            int    `json:"changeInCashAndCashEquivalents,string"`
	ChangeInExchangeRate                                      string `json:"changeInExchangeRate"`
	NetIncome                                                 int    `json:"netIncome,string"`
}

func CashFlowRequest(symbol string) (CashFlow, error) {
	const ENDPOINT_URL string = "CASH_FLOW"

	baseUrl := base_url()
	values := baseUrl.Query()

	values.Add("function", ENDPOINT_URL)
	values.Add("symbol", symbol)

	baseUrl.RawQuery = values.Encode()

	response, err := http.Get(baseUrl.String())

	if err != nil {
		return CashFlow{}, errors.New("the HTTP request is failed with an error")
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		cashFlow := CashFlow{}

		json.Unmarshal(data, &cashFlow)

		return cashFlow, nil
	}
}
