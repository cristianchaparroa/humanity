package websocket

import (
	"fmt"
	"log"

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
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()

		if err != nil {
			log.Println(err)
		}
		message := Message{Type: messageType, Body: string(p)}

		c.Pool.Broadcast <- message
		fmt.Printf("Message received %+v\n", message)

	}
}

// Message contains all the information througth a connection
type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}
