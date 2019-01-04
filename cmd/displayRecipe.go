package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var displayRecipeCommand = &cobra.Command{
	Use:   "display",
	Short: "Display a recipe",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("display called")
	},
}

func init() {
	recipeCmd.AddCommand(displayRecipeCommand)
}
