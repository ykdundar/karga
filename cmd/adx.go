package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/karga/internal"
)

// adxCmd represents the adx command
var adxCmd = &cobra.Command{
	Use:   "adx",
	Short: "This command fetches the average directional movement index (ADX) values.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.ADXRequest(symbol, interval, time_period))
	},
}

func init() {
	rootCmd.AddCommand(adxCmd)
	adxCmd.PersistentFlags().StringVarP(&symbol, "symbol", "s", "IBM", "Stock symbol I.e. IBM.")
	adxCmd.PersistentFlags().StringVarP(&interval, "interval", "i", "weekly", "The following values are supported: 1min, 5min, 15min, 30min, 60min, daily, weekly, monthly.")
	adxCmd.PersistentFlags().IntVarP(&time_period, "timeperiod", "t", 60, "Number of data points used to calculate each RSI value. Positive integers are accepted (e.g., timeperiod=60, timeperiod=200)")

}
