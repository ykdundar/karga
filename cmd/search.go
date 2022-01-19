package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/karga/internal"
)

var keywords string

// symbolCmd represents the symbol command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "This command build an auto-complete search box",
	Long: `This command build an auto-complete search box
	It fetches the best-matching symbols and market information 
	based on keywords of your choice. The search results also contain match scores
	that provide you with the full flexibility to develop your own search and
	filtering logic.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.SearchRequest(symbol, keywords))
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.PersistentFlags().StringVarP(&symbol, "symbol", "s", "IBM", "Stock symbol I.e. IBM.")
	searchCmd.PersistentFlags().StringVarP(&keywords, "keywords", "k", "microsoft", "A text string of your choice. E.g: keywords=microsoft.")
}
