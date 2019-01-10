package actions

import (
	"fmt"
	"os"

	"log"

	"github.com/kdelwat/recipaliser"
	"github.com/kdelwat/recipaliser/db"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var dbPath string

var is recipaliser.IngredientService
var rs recipaliser.RecipeService

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "recipaliser",
	Short: "A recipe creator and database",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	database, err := db.NewDatabase(dbPath)

	if err != nil {
		log.Fatal(err)
	}

	is = &database.IngredientService
	rs = &database.RecipeService
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.recipaliser.yaml)")
	rootCmd.PersistentFlags().StringVarP(&dbPath, "database", "d", "recipaliser.db", "database path (default is recipaliser.db)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".recipaliser")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
