package websocket

// IChatPool defines the methods to use in  the  pool of connections for a chat
// over one room.
type IChatPool interface {

	// Start liste all events in the pool of connections
	Start()

	// Register a client in the pool
	RegisterClient(client IClient) []error

	// Unregister a client in the pool
	UnregisterClient(client IClient) []error

	// Broadcast send a message to all clients in the pool
	BroadcastMessage(message Message) []error

	// GetBroadcastChann retrieves the channel in which is send
	// the messages in the pool
	GetBroadcastChann() chan Message

	// GetUnregisterChann retrieves the channel in wich is
	// unregisterd the clients
	GetUnregisterChann() chan IClient

	// GetRegisterChann retrieves the channel in wich is
	// registed the clients
	GetRegisterChann() chan IClient

	// GetClients retrieves the  clients in the chat pool.
	GetClients() []IClient

	// GetID returns the identification of ChatPool
	GetID() string

	// GetName returns  the name of the chat rooom if is setup by some user.
	GetName() string

	// SetName establish the name of the chat room
	SetName(n string)
}
