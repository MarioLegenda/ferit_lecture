package create

type blogValidationModel struct {
	Title       string `validate:"required"`
	Description string `validate:"required"`
	Content     string `validate:"required"`

	UserID string `validate:"required,uuid4"`
}

func newBlogValidationModel(title string, desc string, content string, userId string) blogValidationModel {
	return blogValidationModel{
		Title:       title,
		Description: desc,
		Content:     content,
		UserID:      userId,
	}
}
