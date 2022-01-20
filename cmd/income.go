package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/karga/internal"
)

// incomeCmd represents the income command
var incomeCmd = &cobra.Command{
	Use:   "income",
	Short: "Annual and quarterly income statements",
	Long: `Annual and quarterly income statements.

Data is generally refreshed on the same day a company reports its latest 
earnings and financials.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.IncomeStatementRequest(symbol))
	},
}

func init() {
	rootCmd.AddCommand(incomeCmd)

	incomeCmd.PersistentFlags().StringVarP(&symbol, "symbol", "s", "IBM", "Stock symbol I.e. IBM.")
}
