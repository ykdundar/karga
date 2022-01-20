package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/karga/internal"
)

// cashflowCmd represents the cashflow command
var cashflowCmd = &cobra.Command{
	Use:   "cashflow",
	Short: "Annual and quarterly cash flow for the company of interest",
	Long: `Annual and quarterly cash flow for the company of interest.
	
Data is generally refreshed on the same day a company reports its latest
earnings and financials.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.CashFlowRequest(symbol))
	},
}

func init() {
	rootCmd.AddCommand(cashflowCmd)

	cashflowCmd.PersistentFlags().StringVarP(&symbol, "symbol", "s", "IBM", "Stock symbol I.e. IBM.")
}
