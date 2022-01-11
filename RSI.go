package main

type RSIValue struct {
	Rsi string `json:"RSI"`
}

type RSIResponse struct {
	MetaData struct {
		OneSymbol          string `json:"1: Symbol"`
		TwoIndicator       string `json:"2: Indicator"`
		ThreeLastRefreshed string `json:"3: Last Refreshed"`
		FourInterval       string `json:"4: Interval"`
		FiveTimePeriod     int    `json:"5: Time Period"`
		SixSeriesType      string `json:"6: Series Type"`
		SevenTimeZone      string `json:"7: Time Zone"`
	} `json:"Meta Data"`
	TechnicalAnalysisRSI map[string]RSIValue `json:"Technical Analysis: RSI"`
}
