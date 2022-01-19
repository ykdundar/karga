package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/karga/internal"
)

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "This command fetches the annual and quarterly balance sheets",
	Long: `This command fetches the annual and quarterly balance sheets 
	for the company of interest, with normalized fields
	mapped to GAAP and IFRS taxonomies of the SEC.

	Data is generally refreshed on the same day a company reports 
	its latest earnings and financials.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.BalanceSheetRequest(symbol))
	},
}

func init() {
	rootCmd.AddCommand(balanceCmd)
	balanceCmd.PersistentFlags().StringVarP(&symbol, "symbol", "s", "IBM", "Stock symbol I.e. IBM.")

}
