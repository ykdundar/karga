package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/karga/internal"
)

// emaCmd represents the ema command
var emaCmd = &cobra.Command{
	Use:   "ema",
	Short: "This command fetches the exponential moving average (EMA) values",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.EMARequest(symbol, interval, time_period, seriesType))
	},
}

func init() {
	rootCmd.AddCommand(emaCmd)

	emaCmd.PersistentFlags().StringVarP(&symbol, "symbol", "s", "IBM", "Stock symbol I.e. IBM.")
	emaCmd.PersistentFlags().StringVarP(&interval, "interval", "i", "weekly", " I.e. 1min, 5min, 15min, 30min, 60min, daily, weekly, monthly.")
	emaCmd.PersistentFlags().IntVarP(&time_period, "timeperiod", "t", 60, "Number of data points used to calculate each RSI value. Positive integers are accepted (e.g., timeperiod=60, timeperiod=200)")
	emaCmd.PersistentFlags().StringVarP(&seriesType, "seriestype", "y", "low", "The desired price type in the time series. Four types are supported: close, open, high, low")

}
