/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/karga/internal"
)

// cashflowCmd represents the cashflow command
var cashflowCmd = &cobra.Command{
	Use:   "cashflow",
	Short: "This command fetches the annual and quarterly cash flow for the company of interest",
	Long: `This command fetches the annual and quarterly cash flow for the company of interest.
	Data is generally refreshed on the same day a company reports its latest earnings and financials.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.CashFlowRequest(symbol))
	},
}

func init() {
	rootCmd.AddCommand(cashflowCmd)
	cashflowCmd.PersistentFlags().StringVarP(&symbol, "symbol", "s", "IBM", "Stock symbol I.e. IBM.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cashflowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cashflowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
