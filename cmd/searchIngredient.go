package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var searchIngredientCommand = &cobra.Command{
	Use:   "search",
	Short: "Search for an ingredient",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
	},
}

func init() {
	ingredientCmd.AddCommand(searchIngredientCommand)
}
