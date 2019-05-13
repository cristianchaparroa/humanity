package websocket

import "github.com/gorilla/websocket"

// IConnection defines the methods to use throught
// a websocket connection
type IConnection interface {
	WriteJSON(o interface{}) error
	ReadJSON(o interface{}) error
	Close() error
}

// Connection is a wrapper  gorilla/websocket.Conn to be
// sure that implements IConnection
type Connection struct {
	// Conn is  pointer to websocket.Conn
	Conn *websocket.Conn
}

// NewConnection generates a pointer to Connection
func NewConnection(c *websocket.Conn) *Connection {
	return &Connection{Conn: c}
}

// WriteJSON writes a json in the websocket connection
func (c *Connection) WriteJSON(o interface{}) error {
	return c.Conn.WriteJSON(o)
}

// ReadJSON reads the data in the websocket connection
func (c *Connection) ReadJSON(o interface{}) error {
	return c.Conn.ReadJSON(o)
}

// Close ends the websocket connection
func (c *Connection) Close() error {
	return c.Conn.Close()
}
