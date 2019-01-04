package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var newRecipeCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new recipe",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("new called")
	},
}

func init() {
	recipeCmd.AddCommand(newRecipeCmd)
}
