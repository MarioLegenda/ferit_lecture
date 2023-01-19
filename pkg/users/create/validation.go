package create

type userValidationModel struct {
	Name     string `validate:"required"`
	LastName string `validate:"required"`
	Email    string `validate:"required,email"`
}

func newUserValidationModel(name string, lastName string, email string) userValidationModel {
	return userValidationModel{
		Name:     name,
		LastName: lastName,
		Email:    email,
	}
}
