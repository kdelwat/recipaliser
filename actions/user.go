package actions

import (
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Perform operations on users",
	Run: func(cmd *cobra.Command, args []string) {
		panic("implement me!")
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
}
