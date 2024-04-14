package uservalidator

const (
	phoneNumberRegex = "^09[0-9]{9}$"
)

type Repository interface {
}
type Validator struct {
	repo Repository
}

func New(repo Repository) Validator {
	return Validator{repo: repo}
}
