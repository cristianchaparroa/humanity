package websocket

import (
	"log"
	"time"

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
	Pool *Pool

	Account *models.Account
}

// Message contains all the information througth a connection
type Message struct {
	ID       string    `json:"id"`
	Type     int       `json:"type"`
	Body     string    `json:"body"`
	UserID   string    `json:"user_id"`
	Nickname string    `json:"nickname"`
	Time     time.Time `json:"time"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		var message Message
		// messageType, p, err := c.Conn.ReadMessage()
		err := c.Conn.ReadJSON(&message)

		message.Type = 1

		if err != nil {
			log.Println(err)
		}
		c.Pool.Broadcast <- message
	}
}
