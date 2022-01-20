package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/karga/internal"
)

var keyword string

// searchCmd represents the symbol command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a stock with a keyword or symbol",
	Long:  "The best-matching symbols and market information based on keyword of your choice.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.SearchRequest(keyword))
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.PersistentFlags().StringVarP(&keyword, "keywords", "k", "microsoft", "A text string of your choice. E.g: microsoft.")
}
