package models

// Account represents the basic
type Account struct {
	ID       string `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	User     *User
}
