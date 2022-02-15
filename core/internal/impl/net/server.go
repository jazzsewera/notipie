package net

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jazzsewera/notipie/core/internal/domain"
	"github.com/jazzsewera/notipie/core/internal/impl/model"
)

func PreflightHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*") // TODO: replace this dev value
	c.Header("Access-Control-Request-Method", "POST, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Connection", "keep-alive")
}

func PushNotificationHandlerFor(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		notification := model.Notification{}
		c.ShouldBindJSON(&notification)
		hub.broadcast <- domainNotificationOf(notification)
	}
}

func domainNotificationOf(n model.Notification) domain.Notification {
	return domain.Notification{}
}

func WSHandlerFor(hub *Hub) gin.HandlerFunc {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // TODO: replace this dev value
		},
	}

	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
		client.hub.register <- client

		go client.writePump()
		go client.readPump()
	}
}