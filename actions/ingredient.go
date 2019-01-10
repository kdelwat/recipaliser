package actions

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ingredientCmd = &cobra.Command{
	Use:   "ingredient",
	Short: "Perform operations on ingredients",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ingredient called")
	},
}

func init() {
	rootCmd.AddCommand(ingredientCmd)
}
