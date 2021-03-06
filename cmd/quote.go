package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/karga/internal"
)

// quoteCmd represents the quote command
var quoteCmd = &cobra.Command{
	Use:   "quote",
	Short: "Price and volume information for a given stock",
	Long: `Price and volume information for a given stock.

For example: Open, High, Low, Price, Volume, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.GlobalQuoteRequest(symbol))
	},
}

func init() {
	rootCmd.AddCommand(quoteCmd)

	quoteCmd.PersistentFlags().StringVarP(&symbol, "symbol", "s", "IBM", "Stock symbol. I.e. IBM")
}
