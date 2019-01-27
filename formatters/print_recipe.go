package formatters

import (
	"fmt"
	"github.com/kdelwat/recipaliser"
	"github.com/olekukonko/tablewriter"
	"os"
)

type recipeNutritionField struct {
	field string
	value float64
}

func PrintRecipe(recipe recipaliser.Recipe) {
	fmt.Printf("Name: %v\n", recipe.Name)

	outputTable := tablewriter.NewWriter(os.Stdout)

	outputTable.SetHeader([]string{"Ingredient", "Amount (g)"})

	for _, ingredient := range recipe.Ingredients {
		outputTable.Append([]string{ingredient.IngredientName, fmt.Sprintf("%v", ingredient.Amount)})
	}

	outputTable.Render()
}

func PrintRecipeNutrition(recipeNutrition recipaliser.RecipeNutrition, selections ...string) {
	outputTable := tablewriter.NewWriter(os.Stdout)

	outputTable.SetHeader([]string{"Nutrient", "Amount (g)"})

	for _, field := range selectNutritionFields(recipeNutrition, selections...) {
		outputTable.Append([]string{field.field, fmt.Sprintf("%f", field.value)})
	}

	outputTable.Render()
}

// TODO: remove this HELLA duplication
func selectNutritionFields(recipeNutrition recipaliser.RecipeNutrition, selections ...string) []recipeNutritionField {
	selectionSets := map[string][]recipeNutritionField{
		"macronutrients": {
			{field: "Energy (kJ)", value: recipeNutrition.EnergyWithDietaryFibre},
			{field: "Available carbohydrates (with sugar alcohols)", value: recipeNutrition.AvailableCarbohydratesWithSugarAlcohols},
			{field: "Dietary fibre", value: recipeNutrition.DietaryFibre},
			{field: "Protein", value: recipeNutrition.Protein},
			{field: "Total fat", value: recipeNutrition.TotalFat},
		},
		"carbohydrates": {
			{field: "Available carbohydrates (with sugar alcohols)", value: recipeNutrition.AvailableCarbohydratesWithSugarAlcohols},
			{field: "Available carbohydrates (without sugar alcohols)", value: recipeNutrition.AvailableCarbohydratesWithoutSugarAlcohol},
			{field: "Starch", value: recipeNutrition.Starch},
			{field: "Total sugars", value: recipeNutrition.TotalSugars},
			{field: "Added sugars", value: recipeNutrition.AddedSugars},
			{field: "Free sugars", value: recipeNutrition.FreeSugars},
		},
		"protein": {
			{field: "Protein", value: recipeNutrition.Protein},
			{field: "Tryptophan", value: recipeNutrition.Tryptophan},
		},
		"fats": {
			{field: "Total fat", value: recipeNutrition.TotalFat},
			{field: "Cholesterol", value: recipeNutrition.Cholesterol},
			{field: "Total saturated fat", value: recipeNutrition.TotalSaturatedFat},
			{field: "Total monounsaturated fat", value: recipeNutrition.TotalMonounsaturatedFat},
			{field: "Total polyunsaturated fat", value: recipeNutrition.TotalPolyunsaturatedFat},
			{field: "Linoleic acid", value: recipeNutrition.LinoleicAcid},
			{field: "Alphalinolenic acid", value: recipeNutrition.AlphalinolenicAcid},
			{field: "EPA", value: recipeNutrition.C205w3Eicosapentaenoic},
			{field: "DPA", value: recipeNutrition.C225w3Docosapentaenoic},
			{field: "DHA", value: recipeNutrition.C226w3Docosahexaenoic},
			{field: "Total long-chain omega-3 fatty acids", value: recipeNutrition.TotalLongChainOmega3FattyAcids},
			{field: "Total trans-fatty acids", value: recipeNutrition.TotalTransFattyAcids},
		},
		"vitamins": {
			{field: "Vitamin A (retinol equivalents)", value: recipeNutrition.VitaminARetinolEquivalents},
			{field: "Thiamin (B1)", value: recipeNutrition.ThiaminB1},
			{field: "Riboflavin (B2)", value: recipeNutrition.RiboflavinB2},
			{field: "Niacin (B3) (derived equivalents)", value: recipeNutrition.NiacinDerivedEquivalents},
			{field: "Dietary folate equivalents", value: recipeNutrition.DietaryFolateEquivalents},
			{field: "Vitamin B6", value: recipeNutrition.VitaminB6},
			{field: "Vitamin B12", value: recipeNutrition.VitaminB12},
			{field: "Vitamin C", value: recipeNutrition.VitaminC},
			{field: "Vitamin E", value: recipeNutrition.VitaminE},
		},
		"minerals": {
			{field: "Calcium (Ca)", value: recipeNutrition.CalciumCa},
			{field: "Iodine (I)", value: recipeNutrition.IodineI},
			{field: "Iron (Fe)", value: recipeNutrition.IronFe},
			{field: "Magnesium (Mg)", value: recipeNutrition.MagnesiumMg},
			{field: "Phosphorus (P)", value: recipeNutrition.PhosphorusP},
			{field: "Potassium (K)", value: recipeNutrition.PotassiumK},
			{field: "Selenium (Se)", value: recipeNutrition.SeleniumSe},
			{field: "Sodium (Na)", value: recipeNutrition.SodiumNa},
			{field: "Zinc (Zn)", value: recipeNutrition.ZincZn},
		},
		"stimulants": {
			{field: "Caffeine", value: recipeNutrition.Caffeine},
		},
		"depressants": {
			{field: "Alcohol", value: recipeNutrition.Alcohol},
		},
	}

	var selectionSet [][]recipeNutritionField
	for _, selection := range selections {
		selectionSet = append(selectionSet, selectionSets[selection])
	}

	return flattenSelectionSetForRecipe(selectionSet)
}

func flattenSelectionSetForRecipe(selectionSet [][]recipeNutritionField) []recipeNutritionField {
	var flattenedFields []recipeNutritionField

	for i := len(selectionSet) - 1; i >= 0; i-- {
		for j := len(selectionSet[i]) - 1; j >= 0; j-- {
			if !containsRecipeNutritionField(flattenedFields, selectionSet[i][j]) {
				flattenedFields = append(flattenedFields, selectionSet[i][j])
			}
		}
	}

	// Reverse fields
	// From https://stackoverflow.com/a/42545484
	for i, j := 0, len(flattenedFields)-1; i < j; i, j = i+1, j-1 {
		flattenedFields[i], flattenedFields[j] = flattenedFields[j], flattenedFields[i]
	}

	return flattenedFields
}

func containsRecipeNutritionField(fields []recipeNutritionField, target recipeNutritionField) bool {
	for i := 0; i < len(fields); i++ {
		if fields[i].field == target.field {
			return true
		}
	}

	return false
}
