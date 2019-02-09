package actions

import (
	"github.com/kdelwat/recipaliser"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var isPregnant bool
var isLactating bool
var bodyFat float64

var newUserCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new user",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		initServices()
		name := args[0]
		age, err := strconv.Atoi(args[1])

		if err != nil {
			log.Fatal(err)
		}

		sex := args[2]

		if sex != "M" && sex != "F" {
			log.Fatal("Sex must be M or F")
		}

		weight, err := strconv.ParseFloat(args[3], 64)

		if err != nil {
			log.Fatal(err)
		}

		if err != nil {
			log.Fatal(err)
		}

		user := recipaliser.User{Name: name, Age: age, Sex: sex, IsPregnant: isPregnant, IsLactating: isLactating, Weight: weight, BodyFatPercentage: bodyFat}

		if err := us.CreateUser(&user); err != nil {
			log.Fatalf("Error: %v", err)
		}
	},
}

func init() {
	newUserCmd.Flags().BoolVar(&isPregnant, "isPregnant", false, "Whether or not the user is pregnant")
	newUserCmd.Flags().BoolVar(&isLactating, "isLactating", false, "Whether or not the user is lactating")
	newUserCmd.Flags().Float64Var(&bodyFat, "bodyfat", -1, "The user's body fat percentage (e.g. 0.12)")

	userCmd.AddCommand(newUserCmd)
}
