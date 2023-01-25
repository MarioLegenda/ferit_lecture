package getting

type userIdValidationModel struct {
	ID string `validate:"required,uuid"`
}

func newUserIdValidationModel(id string) userIdValidationModel {
	return userIdValidationModel{
		ID: id,
	}
}
