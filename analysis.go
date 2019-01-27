package recipaliser

import "reflect"

func CalculateNutritionalData(recipe Recipe, ingredients []Ingredient) (RecipeNutrition, error) {
	var nutrition RecipeNutrition

	for i, recipeIngredient := range recipe.Ingredients {
		adjustedWeightFactor := recipeIngredient.Amount / 100
		nutrition.AdjustedIngredients = append(nutrition.AdjustedIngredients, Ingredient{
			EnergyWithDietaryFibre:                  ingredients[i].EnergyWithDietaryFibre * adjustedWeightFactor,
			EnergyWithoutDietaryFibre:               ingredients[i].EnergyWithoutDietaryFibre * adjustedWeightFactor,
			Moisture:                                ingredients[i].Moisture * adjustedWeightFactor,
			Protein:                                 ingredients[i].Protein * adjustedWeightFactor,
			TotalFat:                                ingredients[i].TotalFat * adjustedWeightFactor,
			AvailableCarbohydratesWithSugarAlcohols: ingredients[i].AvailableCarbohydratesWithSugarAlcohols * adjustedWeightFactor,
			AvailableCarbohydratesWithoutSugarAlcohol: ingredients[i].AvailableCarbohydratesWithoutSugarAlcohol * adjustedWeightFactor,
			Starch:                             ingredients[i].Starch * adjustedWeightFactor,
			TotalSugars:                        ingredients[i].TotalSugars * adjustedWeightFactor,
			AddedSugars:                        ingredients[i].AddedSugars * adjustedWeightFactor,
			FreeSugars:                         ingredients[i].FreeSugars * adjustedWeightFactor,
			DietaryFibre:                       ingredients[i].DietaryFibre * adjustedWeightFactor,
			Alcohol:                            ingredients[i].Alcohol * adjustedWeightFactor,
			Ash:                                ingredients[i].Ash * adjustedWeightFactor,
			PreformedVitaminARetinol:           ingredients[i].PreformedVitaminARetinol * adjustedWeightFactor,
			BetaCarotene:                       ingredients[i].BetaCarotene * adjustedWeightFactor,
			ProvitaminABetaCaroteneEquivalents: ingredients[i].ProvitaminABetaCaroteneEquivalents * adjustedWeightFactor,
			VitaminARetinolEquivalents:         ingredients[i].VitaminARetinolEquivalents * adjustedWeightFactor,
			ThiaminB1:                          ingredients[i].ThiaminB1 * adjustedWeightFactor,
			RiboflavinB2:                       ingredients[i].RiboflavinB2 * adjustedWeightFactor,
			NiacinB3:                           ingredients[i].NiacinB3 * adjustedWeightFactor,
			NiacinDerivedEquivalents:           ingredients[i].NiacinDerivedEquivalents * adjustedWeightFactor,
			FolateNatural:                      ingredients[i].FolateNatural * adjustedWeightFactor,
			FolicAcid:                          ingredients[i].FolicAcid * adjustedWeightFactor,
			TotalFolates:                       ingredients[i].TotalFolates * adjustedWeightFactor,
			DietaryFolateEquivalents:           ingredients[i].DietaryFolateEquivalents * adjustedWeightFactor,
			VitaminB6:                          ingredients[i].VitaminB6 * adjustedWeightFactor,
			VitaminB12:                         ingredients[i].VitaminB12 * adjustedWeightFactor,
			VitaminC:                           ingredients[i].VitaminC * adjustedWeightFactor,
			AlphaTocopherol:                    ingredients[i].AlphaTocopherol * adjustedWeightFactor,
			VitaminE:                           ingredients[i].VitaminE * adjustedWeightFactor,
			CalciumCa:                          ingredients[i].CalciumCa * adjustedWeightFactor,
			IodineI:                            ingredients[i].IodineI * adjustedWeightFactor,
			IronFe:                             ingredients[i].IronFe * adjustedWeightFactor,
			MagnesiumMg:                        ingredients[i].MagnesiumMg * adjustedWeightFactor,
			PhosphorusP:                        ingredients[i].PhosphorusP * adjustedWeightFactor,
			PotassiumK:                         ingredients[i].PotassiumK * adjustedWeightFactor,
			SeleniumSe:                         ingredients[i].SeleniumSe * adjustedWeightFactor,
			SodiumNa:                           ingredients[i].SodiumNa * adjustedWeightFactor,
			ZincZn:                             ingredients[i].ZincZn * adjustedWeightFactor,
			Caffeine:                           ingredients[i].Caffeine * adjustedWeightFactor,
			Cholesterol:                        ingredients[i].Cholesterol * adjustedWeightFactor,
			Tryptophan:                         ingredients[i].Tryptophan * adjustedWeightFactor,
			TotalSaturatedFat:                  ingredients[i].TotalSaturatedFat * adjustedWeightFactor,
			TotalMonounsaturatedFat:            ingredients[i].TotalMonounsaturatedFat * adjustedWeightFactor,
			TotalPolyunsaturatedFat:            ingredients[i].TotalPolyunsaturatedFat * adjustedWeightFactor,
			LinoleicAcid:                       ingredients[i].LinoleicAcid * adjustedWeightFactor,
			AlphalinolenicAcid:                 ingredients[i].AlphalinolenicAcid * adjustedWeightFactor,
			C205w3Eicosapentaenoic:             ingredients[i].C205w3Eicosapentaenoic * adjustedWeightFactor,
			C225w3Docosapentaenoic:             ingredients[i].C225w3Docosapentaenoic * adjustedWeightFactor,
			C226w3Docosahexaenoic:              ingredients[i].C226w3Docosahexaenoic * adjustedWeightFactor,
			TotalLongChainOmega3FattyAcids:     ingredients[i].TotalLongChainOmega3FattyAcids * adjustedWeightFactor,
			TotalTransFattyAcids:               ingredients[i].TotalTransFattyAcids * adjustedWeightFactor,
		})
	}

	nutritionFields := getRecipeNutritionFieldNames(&nutrition)
	nutritionFields = remove(nutritionFields, "AdjustedIngredients")

	for _, nutritionField := range nutritionFields {
		fieldValue := sum(nutrition.AdjustedIngredients, func(i Ingredient) float64 { return getIngredientField(&i, nutritionField) })
		setNutritionField(&nutrition, nutritionField, fieldValue)
	}

	return nutrition, nil
}

func sum(ingredients []Ingredient, accessor func(i Ingredient) float64) float64 {
	s := 0.0
	for _, ingredient := range ingredients {
		s += accessor(ingredient)
	}
	return s
}

// TODO: all uses of reflection are quite unsafe here
// From https://stackoverflow.com/a/18931036
func getIngredientField(ingredient *Ingredient, fieldName string) float64 {
	reflectedIngredient := reflect.ValueOf(ingredient)
	value := reflect.Indirect(reflectedIngredient).FieldByName(fieldName)
	return value.Float()
}

func getRecipeNutritionFieldNames(recipeNutrition *RecipeNutrition) []string {
	reflected := reflect.Indirect(reflect.ValueOf(recipeNutrition))

	numFields := reflected.Type().NumField()

	var fields []string
	for i := 0; i < numFields; i++ {
		fields = append(fields, reflected.Type().Field(i).Name)
	}

	return fields
}

func setNutritionField(recipeNutrition *RecipeNutrition, field string, value float64) {
	reflected := reflect.Indirect(reflect.ValueOf(recipeNutrition))
	reflected.FieldByName(field).SetFloat(value)
}

func remove(strings []string, value string) []string {
	for i, val := range strings {
		if val == value {
			return append(strings[:i], strings[i+1:]...)
		}
	}

	return strings
}
