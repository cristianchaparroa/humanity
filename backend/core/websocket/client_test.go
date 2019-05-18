package websocket

import "testing"

func TestClientWriteMessage(t *testing.T) {

	var test = []struct {
		Message string
	}{
		{"Sending messages"},
		{"This is a test and should sent without errors!"},
		{"Call me, this is urgent!"},
	}

	conn := &WSFackeConn{}

	c := &Client{}
	c.Conn = conn // mock the connection

	for _, tc := range test {
		err := c.WriteMessage(tc.Message)

		if err != nil {
			t.Errorf("The message can't be sent, returns the error:%#v", err)
		}
	}
}
