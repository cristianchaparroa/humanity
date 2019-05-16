package core

import "github.com/streadway/amqp"

// IQueueManager defines the methods to handle the queues
type IQueueManager interface {
	NewConnection(d string) (*amqp.Connection, error)

	ExchangeDeclare(ch *amqp.Channel, name, typ string, autoDel, internal, noWait bool) error

	QueueBind(ch *amqp.Channel, qName, rKey, exchange string) error
}

// QueueManager manages the queues
type QueueManager struct {
}

// NewConnection generates a new connection to specific queue
func NewConnection(d string) (*amqp.Connection, error) {
	return amqp.Dial(d)
}

// ExchangeDeclare creates a exchange to send the messages
func ExchangeDeclare(ch *amqp.Channel, eName, typ string,
	durable, autoDel, internal, noWait bool) error {

	err := ch.ExchangeDeclare(
		eName,    // name
		typ,      // type
		durable,  // durable
		autoDel,  // auto-deleted
		internal, // internal
		noWait,   // no-wait
		nil,      // arguments
	)
	return err
}

// QueueDeclare declares a queue to hold messages and deliver to consumers.
func QueueDeclare(ch *amqp.Channel, qName string, durable, deletUnused, exclusive, noWait bool) (amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		qName,       // name
		durable,     // durable
		deletUnused, // delete when usused
		exclusive,   // exclusive
		noWait,      // no-wait
		nil,         // arguments
	)
	return q, err
}

// QueueBind creates a relation between exchange and queue
func QueueBind(ch *amqp.Channel, qName, rKey, exchange string, noWait bool) error {
	err := ch.QueueBind(
		qName,    // queue name
		rKey,     // routing key
		exchange, // exchange
		noWait,
		nil,
	)
	return err
}

// CreatePublishing setup the publishign
func CreatePublishing(ch *amqp.Channel, cType string, body string) amqp.Publishing {
	p := amqp.Publishing{
		ContentType: cType,
		Body:        []byte(body),
	}
	return p
}

// Publish send a p publication through exchange
func Publish(ch *amqp.Channel, p amqp.Publishing, exchange,
	rKey string, mandatory, inmediate bool) error {

	err := ch.Publish(
		exchange,  // exchange
		rKey,      // routing key
		mandatory, // mandatory
		inmediate, // immediate
		p,         // Publishing
	)
	return err
}
