package websocket

import (
	"net/http/httptest"
	"testing"
)

func TestUpgrade(t *testing.T) {

	w := NewTestResponseWriter()
	r := httptest.NewRequest("GET", "/ws", nil)

	r.Header.Set("Connection", "Upgrade")
	r.Header.Set("Upgrade", "websocket")
	r.Header.Set("Sec-Websocket-Version", "13")
	r.Header.Set("Sec-WebSocket-Key", "websocketsecret-key")

	_, err := Upgrade(w, r)

	if err != nil {
		t.Error(err)
	}
}
