package recipaliser

const (
	Sort_Ascending  = iota
	Sort_Descending = iota
)

// TODO: can embed IngredientService inside here?
type RecipeService interface {
	Recipe(id RecipeID) (Recipe, error)
	RecipeIngredients(id RecipeID, is *IngredientService) ([]Ingredient, error)
	CreateRecipe(recipe *Recipe) error
	AddIngredientToRecipe(id RecipeID, ingredientId IngredientID, amount IngredientAmount, is *IngredientService) error
	RemoveIngredientFromRecipe(id RecipeID, ingredientId IngredientID) error
}

type IngredientService interface {
	Ingredient(id IngredientID) (Ingredient, error)
	CreateIngredient(ingredient *Ingredient) error
	SearchIngredient(nameSubstring string) ([]Ingredient, error)
	ListIngredients(sortField string, sortOrder uint, maxIngredients int) ([]Ingredient, error)
}

type UserService interface {
	User(name UserName) (User, error)
	CreateUser(user *User) error
}

type IngredientID string
type IngredientAmount float64
type Ingredient struct {
	AusnutID                                  string  `db:"ausnut_id"`
	Name                                      string  `db:"name"`
	EnergyWithDietaryFibre                    float64 `db:"energy_with_dietary_fibre"`
	EnergyWithoutDietaryFibre                 float64 `db:"energy_without_dietary_fibre"`
	Moisture                                  float64 `db:"moisture"`
	Protein                                   float64 `db:"protein"`
	TotalFat                                  float64 `db:"total_fat"`
	AvailableCarbohydratesWithSugarAlcohols   float64 `db:"available_carbohydrates_with_sugar_alcohols"`
	AvailableCarbohydratesWithoutSugarAlcohol float64 `db:"available_carbohydrates_without_sugar_alcohol"`
	Starch                                    float64 `db:"starch"`
	TotalSugars                               float64 `db:"total_sugars"`
	AddedSugars                               float64 `db:"added_sugars"`
	FreeSugars                                float64 `db:"free_sugars"`
	DietaryFibre                              float64 `db:"dietary_fibre"`
	Alcohol                                   float64 `db:"alcohol"`
	Ash                                       float64 `db:"ash"`
	PreformedVitaminARetinol                  float64 `db:"preformed_vitamin_a_retinol"`
	BetaCarotene                              float64 `db:"beta_carotene"`
	ProvitaminABetaCaroteneEquivalents        float64 `db:"provitamin_a_beta_carotene_equivalents"`
	VitaminARetinolEquivalents                float64 `db:"vitamin_a_retinol_equivalents"`
	ThiaminB1                                 float64 `db:"thiamin_b1"`
	RiboflavinB2                              float64 `db:"riboflavin_b2"`
	NiacinB3                                  float64 `db:"niacin_b3"`
	NiacinDerivedEquivalents                  float64 `db:"niacin_derived_equivalents"`
	FolateNatural                             float64 `db:"folate_natural"`
	FolicAcid                                 float64 `db:"folic_acid"`
	TotalFolates                              float64 `db:"total_folates"`
	DietaryFolateEquivalents                  float64 `db:"dietary_folate_equivalents"`
	VitaminB6                                 float64 `db:"vitamin_b6"`
	VitaminB12                                float64 `db:"vitamin_b12"`
	VitaminC                                  float64 `db:"vitamin_c"`
	AlphaTocopherol                           float64 `db:"alpha_tocopherol"`
	VitaminE                                  float64 `db:"vitamin_e"`
	CalciumCa                                 float64 `db:"calcium_ca"`
	IodineI                                   float64 `db:"iodine_i"`
	IronFe                                    float64 `db:"iron_fe"`
	MagnesiumMg                               float64 `db:"magnesium_mg"`
	PhosphorusP                               float64 `db:"phosphorus_p"`
	PotassiumK                                float64 `db:"potassium_k"`
	SeleniumSe                                float64 `db:"selenium_se"`
	SodiumNa                                  float64 `db:"sodium_na"`
	ZincZn                                    float64 `db:"zinc_zn"`
	Caffeine                                  float64 `db:"caffeine"`
	Cholesterol                               float64 `db:"cholesterol"`
	Tryptophan                                float64 `db:"tryptophan"`
	TotalSaturatedFat                         float64 `db:"total_saturated_fat"`
	TotalMonounsaturatedFat                   float64 `db:"total_monounsaturated_fat"`
	TotalPolyunsaturatedFat                   float64 `db:"total_polyunsaturated_fat"`
	LinoleicAcid                              float64 `db:"linoleic_acid"`
	AlphalinolenicAcid                        float64 `db:"alphalinolenic_acid"`
	C205w3Eicosapentaenoic                    float64 `db:"c205w3_eicosapentaenoic"`
	C225w3Docosapentaenoic                    float64 `db:"c225w3_docosapentaenoic"`
	C226w3Docosahexaenoic                     float64 `db:"c226w3_docosahexaenoic"`
	TotalLongChainOmega3FattyAcids            float64 `db:"total_long_chain_omega3_fatty_acids"`
	TotalTransFattyAcids                      float64 `db:"total_trans_fatty_acids"`
}

type RecipeID string
type Recipe struct {
	Name        string `db:"name"`
	Ingredients []RecipeIngredient
}

type RecipeIngredient struct {
	RecipeName     string  `db:"recipe_name"`
	IngredientName string  `db:"ingredient_name"`
	Amount         float64 `db:"amount"`
}

// TODO: combine this with the recipe struct
type RecipeNutrition struct {
	AdjustedIngredients                       []Ingredient
	EnergyWithDietaryFibre                    float64
	EnergyWithoutDietaryFibre                 float64
	Moisture                                  float64
	Protein                                   float64
	TotalFat                                  float64
	AvailableCarbohydratesWithSugarAlcohols   float64
	AvailableCarbohydratesWithoutSugarAlcohol float64
	Starch                                    float64
	TotalSugars                               float64
	AddedSugars                               float64
	FreeSugars                                float64
	DietaryFibre                              float64
	Alcohol                                   float64
	Ash                                       float64
	PreformedVitaminARetinol                  float64
	BetaCarotene                              float64
	ProvitaminABetaCaroteneEquivalents        float64
	VitaminARetinolEquivalents                float64
	ThiaminB1                                 float64
	RiboflavinB2                              float64
	NiacinB3                                  float64
	NiacinDerivedEquivalents                  float64
	FolateNatural                             float64
	FolicAcid                                 float64
	TotalFolates                              float64
	DietaryFolateEquivalents                  float64
	VitaminB6                                 float64
	VitaminB12                                float64
	VitaminC                                  float64
	AlphaTocopherol                           float64
	VitaminE                                  float64
	CalciumCa                                 float64
	IodineI                                   float64
	IronFe                                    float64
	MagnesiumMg                               float64
	PhosphorusP                               float64
	PotassiumK                                float64
	SeleniumSe                                float64
	SodiumNa                                  float64
	ZincZn                                    float64
	Caffeine                                  float64
	Cholesterol                               float64
	Tryptophan                                float64
	TotalSaturatedFat                         float64
	TotalMonounsaturatedFat                   float64
	TotalPolyunsaturatedFat                   float64
	LinoleicAcid                              float64
	AlphalinolenicAcid                        float64
	C205w3Eicosapentaenoic                    float64
	C225w3Docosapentaenoic                    float64
	C226w3Docosahexaenoic                     float64
	TotalLongChainOmega3FattyAcids            float64
	TotalTransFattyAcids                      float64
}

type ReferenceValue struct {
	Source string
	Value  float64
}

type ReferenceValues = map[string]ReferenceValue

type UserName = string
type User struct {
	Name                    UserName
	Sex                     string
	Age                     int
	Weight                  float64
	BodyFatPercentage       float64
	NutrientReferenceValues map[string]ReferenceValues
	IsPregnant              bool
	IsLactating             bool
}
