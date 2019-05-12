package websocket

import (
	"log"

	"github.com/cristianchaparroa/humanity/backend/models"
	"github.com/gorilla/websocket"
)

// IClient defines the methods that sould have a client
// in the chat pool
type IClient interface {

	// GetID retrieves the client id
	GetID() string

	// GetPool returns the belongs pool
	// TODO : decouple the pool an replace by IChatPool
	GetPool() *Pool

	//  GetConnection retrieves the websocket connection
	GetConnection() *websocket.Conn

	// Read the message int the pool
	Read()
}

// Client specific to websocket connection.
type Client struct {

	// ID is the string for specific connection
	ID string

	// Conn is  pointer to websocket.Conn
	Conn *websocket.Conn

	// Pool Pointer to the pool which this client will be part
	Pool *Pool

	// Account information related to this client
	Account *models.Account
}

// GetID ...
func (c *Client) GetID() string {
	return c.ID
}

// GetPool ...
func (c *Client) GetPool() *Pool {
	return c.Pool
}

// GetConnection ...
func (c *Client) GetConnection() *websocket.Conn {
	return c.Conn
}

// Read the message in the current connection
func (c *Client) Read() {

	defer func() {
		//pool := c.GetPool()
		//c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		var message Message
		err := c.Conn.ReadJSON(&message)
		message.Type = 1

		if err != nil {
			log.Println(err)
		}
		c.Pool.Broadcast <- message
	}
}
