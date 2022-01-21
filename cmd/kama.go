package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/karga/internal"
)

// kamaCmd represents the kama command
var kamaCmd = &cobra.Command{
	Use:   "kama",
	Short: "Kaufman adaptive moving average (KAMA) values",
	Long:  "Kaufman adaptive moving average (KAMA) values",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.KAMARequest(symbol, interval, timePeriod, seriesType))
	},
}

func init() {
	rootCmd.AddCommand(kamaCmd)

	kamaCmd.PersistentFlags().StringVarP(&symbol, "symbol", "s", "IBM", "Stock symbol I.e. IBM.")
	kamaCmd.PersistentFlags().StringVarP(&interval, "interval", "i", "weekly", " I.e. 1min, 5min, 15min, 30min, 60min, daily, weekly, monthly.")
	kamaCmd.PersistentFlags().IntVarP(&timePeriod, "timeperiod", "t", 60, "Number of data points used to calculate each RSI value. Positive integers are accepted (e.g., timeperiod=60, timeperiod=200)")
	kamaCmd.PersistentFlags().StringVarP(&seriesType, "seriestype", "y", "low", "The desired price type in the time series. Four types are supported: close, open, high, low")
}
