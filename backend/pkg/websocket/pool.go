package websocket

import (
	"fmt"
	"time"
)

// Pool manage the concurrent comunication
type Pool struct {
	// Register channel send to all clients the message
	// that a new client is joined
	Register chan *Client
	// Unregister channel unregister a user and notify the pool
	// when a client is disconnected.
	Unregister chan *Client
	// Cients that are active/inactiv but not disconnected.
	Clients map[*Client]bool
	// Broadcast send a message to all client in the pool
	Broadcast chan Message
}

// NewPool creates a pointer to Pool structure
func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

// Start will be able to listen all events in the pool of connections.
func (p *Pool) Start() {
	for {

		select {
		case client := <-p.Register:
			p.Clients[client] = true
			fmt.Println("Size of connection pool: ", len(p.Clients))

			for c := range p.Clients {
				fmt.Println(client)
				m := Message{Type: 1, Body: "New User joined...", Time: time.Now()}
				c.Conn.WriteJSON(m)
			}

			break
		case client := <-p.Unregister:
			delete(p.Clients, client)
			fmt.Println("Connection pool size: ", len(p.Clients))
			for c := range p.Clients {
				m := Message{Type: 1, Body: "User Disconnected", Time: time.Now()}
				c.Conn.WriteJSON(m)
			}
			break

		case message := <-p.Broadcast:
			fmt.Println("Sending message to all clients in this pool")
			for c := range p.Clients {
				message.Time = time.Now()
				if err := c.Conn.WriteJSON(message); err != nil {

					fmt.Println(err)
					return
				}
			}
			break

		}

	}
}
