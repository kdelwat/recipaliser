package cmd

import (
	"fmt"
	"os"

	"github.com/kdelwat/recipaliser/model"
	"github.com/spf13/cobra"
)

var newRecipeCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new recipe",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := model.NewRecipe(args[0]); err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(0)
		}
	},
}

func init() {
	recipeCmd.AddCommand(newRecipeCmd)
}
