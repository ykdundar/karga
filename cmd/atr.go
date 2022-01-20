package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/karga/internal"
)

// atrCmd represents the atr command
var atrCmd = &cobra.Command{
	Use:   "atr",
	Short: "Average true range (ATR) values",
	Long:  "Average true range (ATR) values",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.ATRRequest(symbol, interval, time_period))
	},
}

func init() {
	rootCmd.AddCommand(atrCmd)

	atrCmd.PersistentFlags().StringVarP(&symbol, "symbol", "s", "IBM", "Stock symbol I.e. IBM.")
	atrCmd.PersistentFlags().StringVarP(&interval, "interval", "i", "weekly", "The following values are supported: 1min, 5min, 15min, 30min, 60min, daily, weekly, monthly.")
	atrCmd.PersistentFlags().IntVarP(&time_period, "timeperiod", "t", 60, "Number of data points used to calculate each RSI value. Positive integers are accepted (e.g., 60, 200)")
}
