package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/karga/internal"
)

var timePeriod int
var seriesType string

var rsiCmd = &cobra.Command{
	Use:   "rsi",
	Short: "This command This fetches the relative strength index (RSI) values.",
	Long:  `This command This fetches the relative strength index (RSI) values.`,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rsiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rsiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
