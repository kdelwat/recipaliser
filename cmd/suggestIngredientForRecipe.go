package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var suggestIngredientForRecipeCommand = &cobra.Command{
	Use:   "suggest",
	Short: "Suggest an ingredient for the recipe",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("suggest called")
	},
}

func init() {
	recipeCmd.AddCommand(suggestIngredientForRecipeCommand)
}
