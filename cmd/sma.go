package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/karga/internal"
)

// smaCmd represents the sma command
var smaCmd = &cobra.Command{
	Use:   "sma",
	Short: "Simple moving average (SMA) values",
	Long:  "Simple moving average (SMA) values",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.SMARequest(symbol, interval, timePeriod, seriesType))
	},
}

func init() {
	rootCmd.AddCommand(smaCmd)

	smaCmd.PersistentFlags().StringVarP(&symbol, "symbol", "s", "IBM", "Stock symbol I.e. IBM.")
	smaCmd.PersistentFlags().StringVarP(&interval, "interval", "i", "weekly", " I.e. 1min, 5min, 15min, 30min, 60min, daily, weekly, monthly.")
	smaCmd.PersistentFlags().IntVarP(&timePeriod, "timeperiod", "t", 60, "Number of data points used to calculate each RSI value. Positive integers are accepted (e.g., timeperiod=60, timeperiod=200)")
	smaCmd.PersistentFlags().StringVarP(&seriesType, "seriestype", "y", "low", "The desired price type in the time series. Four types are supported: close, open, high, low")
}
