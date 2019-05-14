package repositories

import (
	"github.com/cristianchaparroa/humanity/backend/models"
	"github.com/jinzhu/gorm"
)

// IAccountRepository ...
type IAccountRepository interface {
	FindByEmail(email string) (*models.Account, error)
}

// AccountRepository ...
type AccountRepository struct {
	db *gorm.DB
}

// NewAccountRepository generates a pointer to AccountRepository
func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

// FindByEmail search an account by email
func (r *AccountRepository) FindByEmail(email string) (*models.Account, error) {
	//query := `SELECT id,email, password, nickname FROM account WHERE email=$1`
	//row := r.db.QueryRow(query, email)

	var a models.Account
	r.db.Where("email = ?", email).First(&a)

	return &a, nil
}
