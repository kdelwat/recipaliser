package actions

import (
	"fmt"

	"github.com/spf13/cobra"
)

var recipeCmd = &cobra.Command{
	Use:   "recipe",
	Short: "Perform operations on recipes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("recipe called")
	},
}

func init() {
	rootCmd.AddCommand(recipeCmd)
}
