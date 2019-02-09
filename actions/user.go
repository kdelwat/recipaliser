package actions

import (
	"github.com/kdelwat/recipaliser"
	"github.com/kdelwat/recipaliser/formatters"
	"github.com/spf13/cobra"
	"log"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Perform operations on users",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initServices()

		userName := recipaliser.UserName(args[0])

		user, err := us.User(userName)

		if err != nil {
			log.Fatal(err)
		}

		formatters.PrintUser(user)
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
}
