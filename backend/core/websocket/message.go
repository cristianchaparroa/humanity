package websocket

import (
	"time"

	"github.com/google/uuid"
)

// Message contains all the information througth a connection
type Message struct {

	// ID is a uud indetificator for the message
	ID string `json:"id"`

	// Type of message
	Type int `json:"type"`

	// Body contains the information in the message
	Body string `json:"body"`

	// UserID is the user id related to this message
	UserID string `json:"user_id"`

	// Nickname of the user that sends the message
	Nickname string `json:"nickname"`

	// Time in wich was sent the messsage
	Time time.Time `json:"time"`
}

// NewMessage creates a Message
func NewMessage(body string, typ int) Message {
	uuid := uuid.New()
	id := uuid.String()
	t := time.Now()
	m := Message{ID: id, Body: body, Type: typ, Time: t}
	return m
}
