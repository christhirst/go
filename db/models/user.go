package models

type Credentials struct {
	Username string `json:"username", db:"username"`
	password string `json:"password", db:"password"`
}

type UserRepository interface {
	//FindByID(ID int) (*User, error)
	//Save(user *User) error
}
