package app

import (
	"github.com/gin-gonic/gin"
	"github.com/it59com/ws-go-proto/internal/infrastructure/websocket"
	"github.com/it59com/ws-go-proto/internal/interfaces/http"
)

func Run() error {
	r := gin.Default()

	hub := websocket.NewHub()
	go hub.Run()

	handler := http.NewWebSocketHandler(hub)
	r.GET("/ws", handler.HandleWebSocket)

	return r.Run(":8080")
}
