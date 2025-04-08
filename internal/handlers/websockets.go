package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections
		return true
	},
}

func WebsocketRoute(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("ERR:", err)
		return
	}

	for {
		time.Sleep(time.Second * 2)
		err := conn.WriteMessage(
			websocket.TextMessage,
			[]byte(fmt.Sprintf("%s - Live!", time.Now())),
		)
		if err != nil {
			log.Println("ERR:", err)
			return
		}
	}
}
