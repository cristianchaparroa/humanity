package websocket

import (
	"github.com/gorilla/websocket"
)

// IClient defines the methods that sould have a client
// in the chat pool
type IClient interface {

	// GetID retrieves the client id
	GetID() string

	// GetPool returns the belongs pool
	// TODO : decouple the pool an replace by IChatPool
	GetPool() IChatPool

	//  GetConnection retrieves the websocket connection
	GetConnection() *websocket.Conn

	// Read the message int the pool
	Read()
}
