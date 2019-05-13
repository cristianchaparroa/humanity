package websocket

import (
	"testing"

	"github.com/cristianchaparroa/humanity/backend/models"
	"github.com/google/uuid"
)

type WSFackeConn struct{}

func (m *WSFackeConn) WriteJSON(o interface{}) error { return nil }
func (m *WSFackeConn) ReadJSON(o interface{}) error  { return nil }
func (m *WSFackeConn) Close() error                  { return nil }

// GetClients generates a list of clients of size n
func GetClients(num int) []IClient {
	cs := make([]IClient, num)

	for i := 0; i < num; i++ {
		cs = append(cs, &Client{})
	}
	return cs
}

func CreateRandomClient(conn IConnection, cp IChatPool) IClient {
	fakeAcc := models.NewAccount(uuid.New().String(), "e@test.com", "test-nickname")

	return &Client{
		ID:      uuid.New().String(),
		Conn:    conn,
		Pool:    cp,
		Account: fakeAcc,
	}
}

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

func TestChatPoolRegisterClient(t *testing.T) {
	cp := NewChatPool()
	conn := &WSFackeConn{}

	fakeAcc := models.NewAccount("id-test-uuid", "e@test.com", "test-nickname")

	client := &Client{
		ID:      uuid.New().String(),
		Conn:    conn,
		Pool:    cp,
		Account: fakeAcc,
	}

	errs := cp.RegisterClient(client)

	if len(errs) != 0 {
		t.Errorf("It not expected errors but get:%#v: ", errs)
	}
	expectedClients := 1
	resultClients := len(cp.Clients)

	if resultClients != expectedClients {
		t.Errorf("Expected %v clients, but get:%v", expectedClients, resultClients)
	}
}

func TestChatPoolUnregisterClient(t *testing.T) {
	cp := NewChatPool()
	conn := &WSFackeConn{}

	a := CreateRandomClient(conn, cp)
	b := CreateRandomClient(conn, cp)
	cp.RegisterClient(a)
	cp.RegisterClient(b)

	errs := cp.UnregisterClient(a)

	if len(errs) != 0 {
		t.Errorf("It not expected errors but get:%#v: ", errs)
	}

	expectedClients := 1
	resultClients := len(cp.Clients)

	if resultClients != expectedClients {
		t.Errorf("Expected %v clients, but get:%v", expectedClients, resultClients)
	}
}

func TestChatPoolBroadcastMessage(t *testing.T) {
	cp := NewChatPool()
	conn := &WSFackeConn{}

	a := CreateRandomClient(conn, cp)
	b := CreateRandomClient(conn, cp)
	cp.RegisterClient(a)
	cp.RegisterClient(b)

	m := NewMessage("This is a message to broadcast", 1)

	errs := cp.BroadcastMessage(m)

	if len(errs) != 0 {
		t.Errorf("It not expected errors but get:%#v: ", errs)
	}

}
