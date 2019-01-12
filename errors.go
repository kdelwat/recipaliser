package recipaliser

type Error string

const (
	IngredientAlreadyExists = Error("ingredient already exists")
	IngredientNotFound      = Error("ingredient not found")
)

func (e Error) Error() string {
	return string(e)
}
