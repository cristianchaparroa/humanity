package initializer

import (
	"fmt"
	"time"

	"github.com/cristianchaparroa/humanity/backend/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// AccountInitializerService implements the functions to works the first time
type AccountInitializerService struct {
	db *gorm.DB
}

// NewInitialzerService creates a pointer to AccountInitializerService
func NewInitialzerService(db *gorm.DB) *AccountInitializerService {
	return &AccountInitializerService{db: db}
}

// Execute executes all the logic of this initializer
func (s *AccountInitializerService) Execute() error {

	if !s.shouldRun() {
		return nil
	}

	s.CreateTables()
	s.PopulateTables()

	return nil
}

// ShouldInit determintes if the initializer was executed
func (s *AccountInitializerService) shouldRun() bool {

	if s.db.HasTable(&models.Account{}) {
		return false
	}
	return true
}

// CreateTables create all the tables according with the models
func (s *AccountInitializerService) CreateTables() error {

	fmt.Println("--> AccountInitializerService:CreateTables")

	s.db.CreateTable(&models.Account{})

	fmt.Println("<-- AccountInitializerService:CreateTables")
	return nil
}

// PopulateTables fill the tables with the necessary data.
func (s *AccountInitializerService) PopulateTables() error {

	fmt.Println("--> AccountInitializerService:PopulateTables ")
	// Setup the data
	var data = []struct {
		id       string
		email    string
		pass     string
		nickname string
		createAt time.Time
	}{
		{uuid.New().String(), "cristianchaparroa@gmail.com", "12345", "ccchaparroa", time.Now()},
		{uuid.New().String(), "mauriciolopez@gmail.com", "12345", "mlopez", time.Now()},
		{uuid.New().String(), "santiagocastro@gmail.com", "12345", "scastro", time.Now()},
		{uuid.New().String(), "merwinponce@gmail.com", "12345", "mponce", time.Now()},
	}

	// Insert the data
	for _, acc := range data {

		a := &models.Account{ID: acc.id, Email: acc.email, Password: acc.pass,
			Nickname: acc.nickname, CreateAt: acc.createAt}

		s.db.Create(a)
	}

	fmt.Println("<-- AccountInitializerService:PopulateTables ")
	return nil
}
