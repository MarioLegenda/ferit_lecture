package users

type User struct {
	Name     string `form:"name"`
	LastName string `form:"lastName"`
	Email    string `form:"email"`
}
