package websocket

import "testing"

type PoolMock struct{}

func (p *PoolMock) Start()                           {}
func (p *PoolMock) RegisterClient(client IClient)    {}
func (p *PoolMock) UnregisterClient(client IClient)  {}
func (p *PoolMock) BroadcastMessage(message Message) {}
func (p *PoolMock) GetBroadcastChann() chan Message  { return nil }
func (p *PoolMock) GetUnregisterChann() chan IClient { return nil }
func (p *PoolMock) GetRegisterChann() chan IClient   { return nil }
func (p *PoolMock) GetClients() []IClient            { return nil }

func TestNewChatPool(t *testing.T) {
	cp := NewChatPool()

	if cp == nil {
		t.Errorf("Expected a new chat pool but the result is nil")
	}
}

func TestChatPoolGetBroadcastChann(t *testing.T) {
	cp := NewChatPool()

	bchan := cp.GetBroadcastChann()

	if bchan == nil {
		t.Error("Expected a Message channel but get nil ")
	}
}

func TestChatPoolGetUnregisterChann(t *testing.T) {
	cp := NewChatPool()

	uchan := cp.GetUnregisterChann()
	if uchan == nil {
		t.Error("Expected a IClient channel but get nil ")
	}
}

func TestChatPoolGetRegisterChann(t *testing.T) {
	cp := NewChatPool()

	rchan := cp.GetRegisterChann()
	if rchan == nil {
		t.Error("Expected a IClient channel but get nil ")
	}
}

func TestChatPoolGetClients(t *testing.T) {
	var test = []struct {
		Clients         []IClient
		ExpectedClients int
	}{
		{make([]IClient, 0), 0},
		{GetClients(2), 3},
		{GetClients(3), 4},
	}

	for _, tc := range test {
		cp := NewChatPool()
		// add clients to the pool
		for _, client := range tc.Clients {
			cp.Clients[client] = true
		}

		cs := cp.GetClients()

		if len(cs) != tc.ExpectedClients {
			t.Errorf("Expected %v clients, but get:%v", tc.ExpectedClients, len(cs))
		}
	}
}

// GetClients generates a list of clients of size n
func GetClients(num int) []IClient {
	cs := make([]IClient, num)

	for i := 0; i < num; i++ {
		cs = append(cs, &Client{})
	}
	return cs
}
