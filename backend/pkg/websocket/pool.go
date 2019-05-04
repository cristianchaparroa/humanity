package websocket

import (
	"fmt"
	"time"

	"github.com/cristianchaparroa/humanity/backend/services"
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
			RegisterClient(p, client)
			break

		case client := <-p.Unregister:
			UnregisterClient(p, client)
			break

		case message := <-p.Broadcast:
			BroadcastMessage(p, message)
			break

		}

	}
}

// RegisterClient ...
func RegisterClient(p *Pool, client *Client) {
	p.Clients[client] = true
	for c := range p.Clients {
		fmt.Println(client)
		m := Message{Type: 1, Body: fmt.Sprintf(" %s joined... ", client.Account.Nickname), Time: time.Now()}
		c.Conn.WriteJSON(m)
	}

}

// UnregisterClient ...
func UnregisterClient(p *Pool, client *Client) {
	delete(p.Clients, client)
	for c := range p.Clients {
		m := Message{Type: 1, Body: "User Disconnected", Time: time.Now()}
		c.Conn.WriteJSON(m)
	}
}

// BroadcastMessage ...
func BroadcastMessage(p *Pool, message Message) {

	fmt.Println("--> BroadcastMessage Sending message to all clients in this pool")

	ms := make([]Message, 0)
	ms = append(ms, message)

	text := message.Body
	bs := services.NewBotService()
	intent := bs.GetIntent(text)

	if intent == services.StockIntent {
		im := bs.Process(text)

		m := Message{
			Body:     im.RawMessage,
			Nickname: "Bot",
		}
		ms = append(ms, m)
	}

	for c := range p.Clients {

		message.Time = time.Now()

		for _, m := range ms {
			m.Time = time.Now()
			if err := c.Conn.WriteJSON(m); err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	fmt.Println("<-- BroadcastMessage")
}
