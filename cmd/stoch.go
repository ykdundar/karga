package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/karga/internal"
)

// stochCmd represents the stoch command
var stochCmd = &cobra.Command{
	Use:   "stoch",
	Short: "Stochastic oscillator (STOCH) values",
	Long:  "Stochastic oscillator (STOCH) values",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.StochRequest(symbol, interval))
	},
}

func init() {
	rootCmd.AddCommand(stochCmd)

	stochCmd.PersistentFlags().StringVarP(&symbol, "symbol", "s", "IBM", "Stock symbol I.e. IBM.")
	stochCmd.PersistentFlags().StringVarP(&interval, "interval", "i", "1min", "Between two consecutive data points in the time series. I.e. 1min, 5min, 15min, 30min, 60min.")
}
