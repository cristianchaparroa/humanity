package websocket

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// ChatPool implements the behavior to manage the concurrent comunication
// throguth the chat
type ChatPool struct {

	// Register channel send to all clients the message
	// that a new client is joined
	Register chan IClient

	// Unregister channel unregister a user and notify the ChatPool
	// when a client is disconnected.
	Unregister chan IClient

	// Cients that are active/inactiv but not disconnected.
	Clients map[IClient]bool

	// Broadcast send a message to all client in the ChatPool
	Broadcast chan Message
}

// NewChatPool creates a pointer to ChatPool structure
func NewChatPool() *ChatPool {
	return &ChatPool{
		Register:   make(chan IClient),
		Unregister: make(chan IClient),
		Clients:    make(map[IClient]bool),
		Broadcast:  make(chan Message),
	}
}

// Start will be able to listen all events in the ChatPool of connections.
func (p *ChatPool) Start() {
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

// RegisterClient adds a new client in the chat pool
func (p *ChatPool) RegisterClient(client IClient) {

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

// UnregisterClient removes a specific client in the chat pool
func (p *ChatPool) UnregisterClient(client IClient) {

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

// BroadcastMessage send a message to everybody in the chat pool
func (p *ChatPool) BroadcastMessage(m Message) {

	fmt.Println("--> BroadcastMessage Sending message to all clients in this ChatPool")

	for c := range p.Clients {

		m.Time = time.Now()
		m.ID = uuid.New().String()

		fmt.Printf("--> message to send: \n--> %#v \n", m)
		conn := c.GetConnection()

		if err := conn.WriteJSON(m); err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("<-- BroadcastMessage")
}

// GetBroadcastChann retrieves the channel in which is send
// the messages in the ChatPool
func (p *ChatPool) GetBroadcastChann() chan Message {
	return p.Broadcast
}

// GetUnregisterChann retrieves the channel in wich is
// unregisterd the clients
func (p *ChatPool) GetUnregisterChann() chan IClient {
	return p.Unregister
}

// GetRegisterChann retrieves the channel in wich is
// registed the clients
func (p *ChatPool) GetRegisterChann() chan IClient {
	return p.Register
}

// GetClients retrieves the  clients in the chat ChatPool.
func (p *ChatPool) GetClients() []IClient {
	cs := make([]IClient, 0)

	for c := range p.Clients {
		cs = append(cs, c)
	}

	return cs
}
