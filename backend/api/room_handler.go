package api

import (
	"fmt"
	"net/http"

	"github.com/cristianchaparroa/humanity/backend/pkg/websocket"
)

// RoomHandler is in charge to server the web socket connection for
func RoomHandler(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	conn, err := websocket.Upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}
