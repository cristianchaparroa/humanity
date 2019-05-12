package api

import (
	"fmt"
	"net/http"

	"github.com/cristianchaparroa/humanity/backend/core/websocket"
	"github.com/cristianchaparroa/humanity/backend/models"
	"github.com/gin-gonic/gin"
)

// RoomHandler is in charge to server the web socket connection for
func RoomHandler(c *gin.Context, pool websocket.IChatPool, w http.ResponseWriter, r *http.Request) {

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
	registerChann := pool.GetRegisterChann()
	registerChann <- client
	go client.Read()
}
