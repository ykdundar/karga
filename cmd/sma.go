package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/karga/internal"
)

// smaCmd represents the sma command
var smaCmd = &cobra.Command{
	Use:   "sma",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.SMARequest(symbol, interval, time_period, seriesType))
	},
}

func init() {
	rootCmd.AddCommand(smaCmd)
	smaCmd.PersistentFlags().StringVarP(&symbol, "symbol", "s", "IBM", "Stock symbol I.e. IBM.")
	smaCmd.PersistentFlags().StringVarP(&interval, "interval", "i", "weekly", " I.e. 1min, 5min, 15min, 30min, 60min, daily, weekly, monthly.")
	smaCmd.PersistentFlags().IntVarP(&time_period, "timeperiod", "t", 60, "Number of data points used to calculate each RSI value. Positive integers are accepted (e.g., timeperiod=60, timeperiod=200)")
	smaCmd.PersistentFlags().StringVarP(&seriesType, "seriestype", "y", "low", "The desired price type in the time series. Four types are supported: close, open, high, low")

}
