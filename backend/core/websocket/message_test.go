package websocket

import "testing"

func TestNewMessage(t *testing.T) {
	var test = []struct {
		Body string
		Type int
	}{
		{"This is a test to create a message", 1},
		{"Creating another message", 1},
	}

	for _, tc := range test {
		m := NewMessage(tc.Body, tc.Type)

		if len(m.ID) == 0 {
			t.Errorf("The new Message has an empty ID: %#v", m.ID)
		}
	}
}
