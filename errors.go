package recipaliser

type Error string

const (
	IngredientAlreadyExists = Error("ingredient already exists")
	RecipeAlreadyExists     = Error("recipe already exists")
	IngredientNotFound      = Error("ingredient not found")
	RecipeNotFound          = Error("recipe not found")
	NoRelevantUserTemplate  = Error("no relevant user template")
	UserAlreadyExists       = Error("user already exists")
)

func (e Error) Error() string {
	return string(e)
}
