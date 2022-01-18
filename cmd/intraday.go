package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/karga/internal"
)

// intradayCmd represents the intraday command

var interval string
var intradayCmd = &cobra.Command{
	Use:   "intraday",
	Short: "This command returns intraday time series ",
	Long: `This command returns intraday time series of the equity specified,
	covering extended trading hours where applicable
	
	For example:
	
	4:00am to 8:00pm Eastern Time for the US market`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.IntradayRequest(symbol, interval))
	},
}

func init() {
	rootCmd.AddCommand(intradayCmd)

	intradayCmd.PersistentFlags().StringVarP(&symbol, "symbol", "s", "IBM", "Stock symbol I.e. IBM.")
	intradayCmd.PersistentFlags().StringVarP(&interval, "interval", "i", "1min", "Between two consecutive data points in the time series. I.e. 1min, 5min, 15min, 30min, 60min.")

}
