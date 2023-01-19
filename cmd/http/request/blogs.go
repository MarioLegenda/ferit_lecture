package request

type Blog struct {
	UserID      string `form:"userId"`
	Title       string `form:"title"`
	Content     string `form:"content"`
	Description string `form:"description"`
}
