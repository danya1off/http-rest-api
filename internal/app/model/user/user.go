package user

// DAO ...
type DAO interface {
	Create(u *User) (*User, error)
	FindByEmail(email string) (*User, error)
}

// User ...
type User struct {
	ID        int
	Email     string
	EncPasswd string
}
