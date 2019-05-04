package api

import (
	"fmt"
	"net/http"

	"github.com/cristianchaparroa/humanity/backend/models"
	"github.com/cristianchaparroa/humanity/backend/pkg/websocket"
	"github.com/gin-gonic/gin"
)

// RoomHandler is in charge to server the web socket connection for
func RoomHandler(c *gin.Context, pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("RoomHandler")
	fmt.Println(r.Host)

	userID := c.Query("userId")
	nickname := c.Query("nickname")

	acc := models.NewAccount(userID, "", nickname)

	conn, err := websocket.Upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn:    conn,
		Pool:    pool,
		Account: acc,
	}

	pool.Register <- client
	go client.Read()
}
