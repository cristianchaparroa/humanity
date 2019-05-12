package websocket

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// IChatPool defines the methods to use in  the  pool of connections for a chat
// over one room.
type IChatPool interface {

	// Start liste all events in the pool of connections
	Start()

	// Register a client in the pool
	RegisterClient(client IClient)

	// Unregister a client in the pool
	UnregisterClient(client IClient)

	// Broadcast send a message to all clients in the pool
	BroadcastMessage(message Message)
}

// Pool manage the concurrent comunication
type Pool struct {

	// Register channel send to all clients the message
	// that a new client is joined
	Register chan IClient

	// Unregister channel unregister a user and notify the pool
	// when a client is disconnected.
	Unregister chan IClient

	// Cients that are active/inactiv but not disconnected.
	Clients map[IClient]bool

	// Broadcast send a message to all client in the pool
	Broadcast chan Message
}

// NewPool creates a pointer to Pool structure
func NewPool() *Pool {
	return &Pool{
		Register:   make(chan IClient),
		Unregister: make(chan IClient),
		Clients:    make(map[IClient]bool),
		Broadcast:  make(chan Message),
	}
}

// Start will be able to listen all events in the pool of connections.
func (p *Pool) Start() {
	for {

		select {
		case client := <-p.Register:
			p.RegisterClient(client)
			break

		case client := <-p.Unregister:
			p.UnregisterClient(client)
			break

		case message := <-p.Broadcast:
			p.BroadcastMessage(message)
			break
		}
	}
}

// RegisterClient ...
func (p *Pool) RegisterClient(client IClient) {

	fmt.Println("--> RegisterClient")

	p.Clients[client] = true

	for c := range p.Clients {
		body := fmt.Sprintf(" %s joined... ", "")
		m := NewMessage(body, 1)
		conn := c.GetConnection()
		conn.WriteJSON(m)
	}

	fmt.Println("<-- RegisterClient")
}

// UnregisterClient ...
func (p *Pool) UnregisterClient(client IClient) {

	fmt.Println("--> UnregisterClient")

	delete(p.Clients, client)
	for c := range p.Clients {
		body := "User Disconnected"
		m := NewMessage(body, 1)
		conn := c.GetConnection()
		conn.WriteJSON(m)
	}

	fmt.Println("<-- UnregisterClient")
}

// BroadcastMessage ...
func (p *Pool) BroadcastMessage(m Message) {

	fmt.Println("--> BroadcastMessage Sending message to all clients in this pool")

	for c := range p.Clients {

		uuid := uuid.New()
		m.Time = time.Now()
		m.ID = uuid.String()

		fmt.Printf("--> message to send: \n--> %#v \n", m)
		conn := c.GetConnection()

		if err := conn.WriteJSON(m); err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("<-- BroadcastMessage")
}
