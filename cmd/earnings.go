package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/karga/internal"
)

// earningsCmd represents the earnings command
var earningsCmd = &cobra.Command{
	Use:   "earnings",
	Short: "This command fetches the annual and quarterly earnings for the company of interest.",
	Long: `This command fetches the annual and quarterly earnings (EPS) for the company of interest.
	Quarterly data also includes analyst estimates and surprise metrics.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.EarningsRequest(symbol))
	},
}

func init() {
	rootCmd.AddCommand(earningsCmd)
	earningsCmd.PersistentFlags().StringVarP(&symbol, "symbol", "s", "IBM", "Stock symbol I.e. IBM.")
}
