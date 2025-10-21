package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/it59com/ws-go-proto/internal/infrastructure/websocket"
)

type WebSocketHandler struct {
	hub *websocket.Hub
}

func NewWebSocketHandler(hub *websocket.Hub) *WebSocketHandler {
	return &WebSocketHandler{hub: hub}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *gin.Context) bool { return true },
}

func (h *WebSocketHandler) HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := websocket.NewClient(h.hub, conn)
	h.hub.register <- client

	go client.WritePump()
	go client.ReadPump()
}
