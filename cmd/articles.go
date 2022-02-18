/*
Copyright © 2022 Harjyot Bagga baggaharjyot@gmail.com

*/
package cmd

import (
	public_operations "staq/pkg/stack_exchange/public"

	"github.com/spf13/cobra"
)

// articlesCmd represents the articles command
var articlesCmd = &cobra.Command{
	Use:   "articles",
	Short: "Get StackExchange sites.",
	Long:  `Get StackExchange sites.`,
	Run: func(cmd *cobra.Command, args []string) {
		public_operations.GetStackExchangeArticles()
	},
}

func init() {
	rootCmd.AddCommand(articlesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sitesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sitesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
