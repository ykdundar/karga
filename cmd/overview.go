package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/karga/internal"
)

// overviewCmd represents the overview command
var overviewCmd = &cobra.Command{
	Use:   "overview",
	Short: "Overview of related company information",
	Long:  "Company information, financial ratios and other key metrics for the equity specified.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.OverviewRequest(symbol))
	},
}

func init() {
	rootCmd.AddCommand(overviewCmd)

	overviewCmd.PersistentFlags().StringVarP(&symbol, "symbol", "s", "IBM", "Stock symbol I.e. IBM.")
}
