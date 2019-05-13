package websocket

import (
	"github.com/cristianchaparroa/humanity/backend/models"
)

// IClient defines the methods that sould have a client
// in the chat pool
type IClient interface {

	// GetID retrieves the client id
	GetID() string

	//  GetConnection retrieves the websocket connection
	GetConnection() IConnection

	// GetPool returns the belongs pool
	GetPool() IChatPool

	// GetUser returns the information related to user that
	// uses the client connection
	GetUser() *models.Account

	// Read the message int the pool
	Read()

	// WriteMessage writes a message in the pool
	WriteMessage(m interface{}) error
}
