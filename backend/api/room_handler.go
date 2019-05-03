package api

import (
	"fmt"

	"github.com/cristianchaparroa/humanity/backend/pkg/websocket"
	"github.com/gin-gonic/gin"
)

// RoomHandler is in charge to server the web socket connection for
func RoomHandler(pool *websocket.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		w := c.Writer
		r := c.Request
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

	return gin.HandlerFunc(fn)
}
