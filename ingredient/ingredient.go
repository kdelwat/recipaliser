package ingredient

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/kdelwat/recipaliser/db"
)

type Ingredient struct {
	gorm.Model
	AusnutID                                  string
	Name                                      string
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

func New(ingredient Ingredient) (Ingredient, error) {
	var existingIngredient Ingredient

	err := db.Db.First(&existingIngredient, "ausnut_id = ?", ingredient.AusnutID).Error

	if err != gorm.ErrRecordNotFound {
		return ingredient, errors.New("an ingredient with the same AUSNUT ID already exists")
	}

	err = db.Db.Create(&ingredient).Error

	if err != nil {
		return ingredient, errors.New(fmt.Sprintf("Could not create new ingredient in database: %v", err))
	}

	return ingredient, nil

}

func Search(searchString string) ([]Ingredient, error) {
	var matchingIngredients []Ingredient

	err := db.Db.Where("name LIKE ?", "%"+searchString+"%").Find(&matchingIngredients).Error

	return matchingIngredients, err
}
