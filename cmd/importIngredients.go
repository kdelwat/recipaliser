package cmd

import (
	"fmt"

	"bufio"
	"encoding/csv"
	"io"
	"os"

	"strconv"

	"github.com/kdelwat/recipaliser/ingredient"
	"github.com/spf13/cobra"
	"gopkg.in/cheggaaa/pb.v1"
)

var importIngredientsCommand = &cobra.Command{
	Use:   "import",
	Short: "Import ingredients into the database",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		importIngredients(args[0])
	},
}

func init() {
	ingredientCmd.AddCommand(importIngredientsCommand)
}

func importIngredients(filename string) {
	csvFile, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Could not import %v: %v", filename, err)
		os.Exit(1)
	}

	csvReader := csv.NewReader(bufio.NewReader(csvFile))

	_, err = csvReader.Read()

	if err != nil {
		fmt.Printf("Could not read header line of %v: %v", filename, err)
	}

	var ingredients []ingredient.Ingredient

	lineNo := 1
	for {
		line, err := csvReader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Could not import %v at line %v: %v", filename, lineNo, err)
			os.Exit(1)
		}

		var nutritionalData []float64
		for _, nutritionalInfo := range line[4:] {
			nutritionalInfoAsFloat, err := strconv.ParseFloat(nutritionalInfo, 64)

			if err != nil {
				fmt.Printf("Could not import %v at line %v: could not convert %v to float", filename, lineNo, nutritionalInfo)
				os.Exit(1)

			}

			nutritionalData = append(nutritionalData, nutritionalInfoAsFloat)
		}

		ingredients = append(ingredients, ingredient.Ingredient{
			AusnutID: line[0],
			Name:     line[2],
			EnergyWithDietaryFibre:    nutritionalData[0],
			EnergyWithoutDietaryFibre: nutritionalData[1],
			Moisture:                  nutritionalData[2],
			Protein:                   nutritionalData[3],
			TotalFat:                  nutritionalData[4],
			AvailableCarbohydratesWithSugarAlcohols:   nutritionalData[5],
			AvailableCarbohydratesWithoutSugarAlcohol: nutritionalData[6],
			Starch:       nutritionalData[7],
			TotalSugars:  nutritionalData[8],
			AddedSugars:  nutritionalData[9],
			FreeSugars:   nutritionalData[10],
			DietaryFibre: nutritionalData[11],
			Alcohol:      nutritionalData[12],
			Ash:          nutritionalData[13],
			PreformedVitaminARetinol:           nutritionalData[14],
			BetaCarotene:                       nutritionalData[15],
			ProvitaminABetaCaroteneEquivalents: nutritionalData[16],
			VitaminARetinolEquivalents:         nutritionalData[17],
			ThiaminB1:                          nutritionalData[18],
			RiboflavinB2:                       nutritionalData[19],
			NiacinB3:                           nutritionalData[20],
			NiacinDerivedEquivalents:           nutritionalData[21],
			FolateNatural:                      nutritionalData[22],
			FolicAcid:                          nutritionalData[23],
			TotalFolates:                       nutritionalData[24],
			DietaryFolateEquivalents:           nutritionalData[25],
			VitaminB6:                          nutritionalData[26],
			VitaminB12:                         nutritionalData[27],
			VitaminC:                           nutritionalData[28],
			AlphaTocopherol:                    nutritionalData[29],
			VitaminE:                           nutritionalData[30],
			CalciumCa:                          nutritionalData[31],
			IodineI:                            nutritionalData[32],
			IronFe:                             nutritionalData[33],
			MagnesiumMg:                        nutritionalData[34],
			PhosphorusP:                        nutritionalData[35],
			PotassiumK:                         nutritionalData[36],
			SeleniumSe:                         nutritionalData[37],
			SodiumNa:                           nutritionalData[38],
			ZincZn:                             nutritionalData[39],
			Caffeine:                           nutritionalData[40],
			Cholesterol:                        nutritionalData[41],
			Tryptophan:                         nutritionalData[42],
			TotalSaturatedFat:                  nutritionalData[43],
			TotalMonounsaturatedFat:            nutritionalData[44],
			TotalPolyunsaturatedFat:            nutritionalData[45],
			LinoleicAcid:                       nutritionalData[46],
			AlphalinolenicAcid:                 nutritionalData[47],
			C205w3Eicosapentaenoic:             nutritionalData[48],
			C225w3Docosapentaenoic:             nutritionalData[49],
			C226w3Docosahexaenoic:              nutritionalData[50],
			TotalLongChainOmega3FattyAcids:     nutritionalData[51],
			TotalTransFattyAcids:               nutritionalData[52]})

		lineNo++
	}

	progressBar := pb.StartNew(len(ingredients))

	ingredientsImported := 0
	ingredientsSkipped := 0
	for _, i := range ingredients {
		_, err = ingredient.New(i)

		if err.Error() == "an ingredient with the same AUSNUT ID already exists" {
			ingredientsSkipped++
		} else if err != nil {
			fmt.Printf("Could not import %v: %v", filename, err)
			os.Exit(1)
		} else {
			ingredientsImported++
		}

		progressBar.Increment()
	}

	progressBar.FinishPrint("Complete")

	fmt.Println("Import completed successfully")
	fmt.Printf("Imported %v ingredients\n", ingredientsImported)
	fmt.Printf("Skipped %v ingredients that were already present in database\n", ingredientsSkipped)
}
