package websocket

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

// ChatPool implements the behavior to manage the concurrent comunication
// throguth the chat
type ChatPool struct {
	ID string

	Name string

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

	// TODO: check how to decompose this implementation.
	MQConn *amqp.Connection
}

// NewChatPool creates a pointer to ChatPool structure
func NewChatPool() *ChatPool {
	return &ChatPool{
		ID:         uuid.New().String(),
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
func (p *ChatPool) RegisterClient(client IClient) []error {

	fmt.Println("--> RegisterClient")
	es := make([]error, 0)

	p.Clients[client] = true

	for c := range p.Clients {
		name := c.GetUser().Nickname
		body := fmt.Sprintf(" %s joined... ", name)
		m := NewMessage(body, 1)
		err := c.WriteMessage(m)

		if err != nil {
			es = append(es, err)
		}
	}

	fmt.Println("<-- RegisterClient")

	return es
}

// UnregisterClient removes a specific client in the chat pool
func (p *ChatPool) UnregisterClient(client IClient) []error {

	fmt.Println("--> UnregisterClient")
	es := make([]error, 0)

	delete(p.Clients, client)

	for c := range p.Clients {
		name := c.GetUser().Nickname
		body := fmt.Sprintf("%s user Disconnected", name)
		m := NewMessage(body, 1)
		err := c.WriteMessage(m)

		if err != nil {
			es = append(es, err)
		}
	}

	fmt.Println("<-- UnregisterClient")
	return nil
}

// BroadcastMessage send a message to everybody in the chat pool
func (p *ChatPool) BroadcastMessage(m Message) []error {

	fmt.Println("--> BroadcastMessage Sending message to all clients in this ChatPool")

	es := make([]error, 0)

	for c := range p.Clients {

		m.Time = time.Now()
		m.ID = uuid.New().String()
		if err := c.WriteMessage(m); err != nil {
			es = append(es, err)
		}
	}
	fmt.Println("<-- BroadcastMessage")
	return es
}

// GetBroadcastChann retrieves the channel in which is send
// the messages in the ChatPool
func (p *ChatPool) GetBroadcastChann() chan Message {
	return p.Broadcast
}

// GetUnregisterChann retrieves the channel in which is
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

// GetID retrieves the ChatPool identificator
func (p *ChatPool) GetID() string {
	return p.ID
}

// GetName get the name of the current pool
func (p *ChatPool) GetName() string {
	return p.Name
}

// SetName set a new name for the current pool
func (p *ChatPool) SetName(n string) {
	p.Name = n
}
