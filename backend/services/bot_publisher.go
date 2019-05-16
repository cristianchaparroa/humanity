package services

import "github.com/cristianchaparroa/humanity/backend/core/websocket"

// IBotPublisher is in charge to defines the methods to
// operate with queues.
type IBotPublisher interface {
	Publish(m interface{}, p websocket.IChatPool) error
}

// BotPublisher ois in charge to implements the mehtods to
// operate with queues.
type BotPublisher struct {
}

//NewBotPublisher returns a pointer to BotPublisher
func NewBotPublisher() *BotPublisher {
	return &BotPublisher{}
}

// Publish sent a message into the queue.
func (b *BotPublisher) Publish(m interface{}, p websocket.IChatPool) error {
	return nil
}
