package repositories

import (
	"database/sql"
	"fmt"

	"github.com/cristianchaparroa/humanity/backend/models"
)

// IAccountRepository ...
type IAccountRepository interface {
	FindByEmail(email string) (*models.Account, error)
}

// AccountRepository ...
type AccountRepository struct {
	db *sql.DB
}

// NewAccountRepository generates a pointer to AccountRepository
func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

// FindByEmail search an account by email
func (r *AccountRepository) FindByEmail(email string) (*models.Account, error) {
	query := `SELECT id,email, password FROM account WHERE email=$1`
	row := r.db.QueryRow(query, email)

	a := &models.Account{}

	err := row.Scan(&a.ID, &a.Email, &a.Password)

	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	return a, nil
}
