package entities

type User struct {
	Name string  // public
	email string  // private
}

type user struct {
	Name string
	Email string
}

type Admin struct {
	user  // private
	Rights int  // public
}

