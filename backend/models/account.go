package models

// Account represents the basic
type Account struct {
	ID       string `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Nickname string `db:"nickname"`
}

// NewAccount generates a pointer to Account
func NewAccount(id, email, nickname string) *Account {
	return &Account{
		ID:       id,
		Email:    email,
		Nickname: nickname,
	}
}
