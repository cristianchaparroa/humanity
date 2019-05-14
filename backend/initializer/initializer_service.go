package initializer

import "github.com/jinzhu/gorm"

// InitializerService defines the operations that sould be performed the first time
// that start the application.
type InitializerService interface {

	//Execute
	Execute() error
}

// InitialzerManager is in charge to execute all initializer
type InitialzerManager struct {
}

func NewInitialzerManager() *InitialzerManager {
	return &InitialzerManager{}
}

// Run  should register all initializer and execute thems
func (m *InitialzerManager) Run(db *gorm.DB) {
	initializers := make([]InitializerService, 0)
	initializers = append(initializers, NewInitialzerService(db))

	for _, init := range initializers {
		init.Execute()
	}
}
