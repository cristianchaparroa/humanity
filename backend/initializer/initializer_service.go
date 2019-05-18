package initializer

import "github.com/jinzhu/gorm"

// InitService defines the operations that sould be performed the first time
// that start the application.
type InitService interface {

	//Execute
	Execute() error
}

// InitialzerManager is in charge to execute all initializer
type InitialzerManager struct {
}

// NewInitialzerManager generates  a pointer to InitialzerManager
func NewInitialzerManager() *InitialzerManager {
	return &InitialzerManager{}
}

// Run  should register all initializer and execute thems
func (m *InitialzerManager) Run(db *gorm.DB) {
	initializers := make([]InitService, 0)

	// Add all the initializer here.
	initializers = append(initializers, NewAccInitialzerService(db))

	for _, init := range initializers {
		init.Execute()
	}
}
