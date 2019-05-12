package websocket

import (
	"log"

	"github.com/cristianchaparroa/humanity/backend/models"
	"github.com/gorilla/websocket"
)

// Client specific to websocket connection.
type Client struct {

	// ID is the string for specific connection
	ID string

	// Conn is  pointer to websocket.Conn
	Conn *websocket.Conn

	// Pool Pointer to the pool which this client will be part
	Pool IChatPool

	// Account information related to this client
	Account *models.Account
}

// Read the message in the current connection
func (c *Client) Read() {

	defer func() {
		pool := c.GetPool()

		uchann := pool.GetUnregisterChann()
		uchann <- c
		c.Conn.Close()
	}()

	for {
		var message Message
		err := c.Conn.ReadJSON(&message)

		if err != nil {
			log.Println(err)
		}
		bchann := c.Pool.GetBroadcastChann()
		bchann <- message
	}
}

// WriteMessage writes a message in the pool
func (c *Client) WriteMessage(m interface{}) error {
	conn := c.GetConnection()
	err := conn.WriteJSON(m)
	return err
}

// GetID retrieves the client id
func (c *Client) GetID() string {
	return c.ID
}

// GetPool returns the chat pool in wich is connected
func (c *Client) GetPool() IChatPool {
	return c.Pool
}

// GetConnection  returns the web socket connection.
func (c *Client) GetConnection() *websocket.Conn {
	return c.Conn
}

// GetUser returns the information related to user that
// uses the client connection
func (c *Client) GetUser() *models.Account {
	return c.Account
}
