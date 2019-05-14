package services

import (
	"github.com/cristianchaparroa/humanity/backend/models"
	"github.com/cristianchaparroa/humanity/backend/repositories"
	"github.com/jinzhu/gorm"
)

// IAccountService defines the services related with accounts
type IAccountService interface {
	Login(email, password string) (bool, string)
}

// AccountService implemtents ILoginService
type AccountService struct {
	db *gorm.DB
}

// NewAccountService generates a pointer to LoginService
func NewAccountService(db *gorm.DB) *AccountService {
	return &AccountService{db: db}
}

// Login verifies if the email and password are the rigth.
func (s *AccountService) Login(email, password string) (bool, *models.Account) {
	ar := repositories.NewAccountRepository(s.db)
	account, err := ar.FindByEmail(email)

	if err != nil {
		return false, nil
	}

	if password == account.Password {
		return true, account
	}

	return false, nil
}
