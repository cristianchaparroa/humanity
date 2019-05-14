package models

import "time"

// Account represents the basic
type Account struct {
	ID       string `gorm:"primary_key"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Nickname string `db:"nickname"`
	CreateAt time.Time
}

// NewAccount generates a pointer to Account
func NewAccount(id, email, nickname string) *Account {
	return &Account{
		ID:       id,
		Email:    email,
		Nickname: nickname,
	}
}
