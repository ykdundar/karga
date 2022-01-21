package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/karga/internal"
)

var rsiCmd = &cobra.Command{
	Use:   "rsi",
	Short: "Relative strength index (RSI) values",
	Long:  "Relative strength index (RSI) values",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.RSIRequest(symbol, interval, timePeriod, seriesType))
	},
}

func init() {
	rootCmd.AddCommand(rsiCmd)

	rsiCmd.PersistentFlags().StringVarP(&symbol, "symbol", "s", "IBM", "Stock symbol I.e. IBM.")
	rsiCmd.PersistentFlags().StringVarP(&interval, "interval", "i", "weekly", " I.e. 1min, 5min, 15min, 30min, 60min, daily, weekly, monthly.")
	rsiCmd.PersistentFlags().IntVarP(&timePeriod, "timeperiod", "t", 60, "Number of data points used to calculate each RSI value. Positive integers are accepted (e.g., timeperiod=60, timeperiod=200)")
	rsiCmd.PersistentFlags().StringVarP(&seriesType, "seriestype", "y", "low", "The desired price type in the time series. Four types are supported: close, open, high, low")
}
