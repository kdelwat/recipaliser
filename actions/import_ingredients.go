package actions

import (
	"fmt"

	"bufio"
	"encoding/csv"
	"io"
	"os"

	"strconv"

	"github.com/kdelwat/recipaliser"
	"github.com/spf13/cobra"
	"gopkg.in/cheggaaa/pb.v1"
)

var importIngredientsCommand = &cobra.Command{
	Use:   "import",
	Short: "Import ingredients into the database",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initServices()
		importIngredients(args[0])
	},
}

func init() {
	ingredientCmd.AddCommand(importIngredientsCommand)
}

func ugToG(n float64) float64 {
	return n / 1000000
}

func mgToG(n float64) float64 {
	return n / 1000
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

	var ingredients []recipaliser.Ingredient

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

		ingredients = append(ingredients, recipaliser.Ingredient{
			AusnutID:                                line[0],
			Name:                                    line[2],
			EnergyWithDietaryFibre:                  nutritionalData[0],
			EnergyWithoutDietaryFibre:               nutritionalData[1],
			Moisture:                                nutritionalData[2],
			Protein:                                 nutritionalData[3],
			TotalFat:                                nutritionalData[4],
			AvailableCarbohydratesWithSugarAlcohols: nutritionalData[5],
			AvailableCarbohydratesWithoutSugarAlcohol: nutritionalData[6],
			Starch:                             nutritionalData[7],
			TotalSugars:                        nutritionalData[8],
			AddedSugars:                        nutritionalData[9],
			FreeSugars:                         nutritionalData[10],
			DietaryFibre:                       nutritionalData[11],
			Alcohol:                            nutritionalData[12],
			Ash:                                nutritionalData[13],
			PreformedVitaminARetinol:           ugToG(nutritionalData[14]),
			BetaCarotene:                       ugToG(nutritionalData[15]),
			ProvitaminABetaCaroteneEquivalents: ugToG(nutritionalData[16]),
			VitaminARetinolEquivalents:         ugToG(nutritionalData[17]),
			ThiaminB1:                          mgToG(nutritionalData[18]),
			RiboflavinB2:                       mgToG(nutritionalData[19]),
			NiacinB3:                           mgToG(nutritionalData[20]),
			NiacinDerivedEquivalents:           mgToG(nutritionalData[21]),
			FolateNatural:                      ugToG(nutritionalData[22]),
			FolicAcid:                          ugToG(nutritionalData[23]),
			TotalFolates:                       ugToG(nutritionalData[24]),
			DietaryFolateEquivalents:           ugToG(nutritionalData[25]),
			VitaminB6:                          mgToG(nutritionalData[26]),
			VitaminB12:                         ugToG(nutritionalData[27]),
			VitaminC:                           mgToG(nutritionalData[28]),
			AlphaTocopherol:                    mgToG(nutritionalData[29]),
			VitaminE:                           mgToG(nutritionalData[30]),
			CalciumCa:                          mgToG(nutritionalData[31]),
			IodineI:                            ugToG(nutritionalData[32]),
			IronFe:                             mgToG(nutritionalData[33]),
			MagnesiumMg:                        mgToG(nutritionalData[34]),
			PhosphorusP:                        mgToG(nutritionalData[35]),
			PotassiumK:                         mgToG(nutritionalData[36]),
			SeleniumSe:                         ugToG(nutritionalData[37]),
			SodiumNa:                           mgToG(nutritionalData[38]),
			ZincZn:                             mgToG(nutritionalData[39]),
			Caffeine:                           mgToG(nutritionalData[40]),
			Cholesterol:                        mgToG(nutritionalData[41]),
			Tryptophan:                         mgToG(nutritionalData[42]),
			TotalSaturatedFat:                  nutritionalData[43],
			TotalMonounsaturatedFat:            nutritionalData[44],
			TotalPolyunsaturatedFat:            nutritionalData[45],
			LinoleicAcid:                       nutritionalData[46],
			AlphalinolenicAcid:                 nutritionalData[47],
			C205w3Eicosapentaenoic:             mgToG(nutritionalData[48]),
			C225w3Docosapentaenoic:             mgToG(nutritionalData[49]),
			C226w3Docosahexaenoic:              mgToG(nutritionalData[50]),
			TotalLongChainOmega3FattyAcids:     mgToG(nutritionalData[51]),
			TotalTransFattyAcids:               mgToG(nutritionalData[52])})

		lineNo++
	}

	progressBar := pb.StartNew(len(ingredients))

	ingredientsImported := 0
	ingredientsSkipped := 0
	for _, i := range ingredients {
		err = is.CreateIngredient(&i)

		if err == recipaliser.IngredientAlreadyExists {
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
