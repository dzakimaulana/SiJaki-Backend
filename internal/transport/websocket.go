package transport

import (
	"log"

	"github.com/gofiber/websocket/v2"
)

func WebSocketHandler(c *websocket.Conn) {
	log.Println("WebSocket connection established")

	if err := sendMessage(c, websocket.TextMessage, []byte("Welcome to the WebSocket server!")); err != nil {
		log.Println("Error sending welcome message:", err)
		return
	}

	listenMessage(c)
}

func listenMessage(c *websocket.Conn) {
	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		log.Printf("Received message: %s", msg)

		if err := sendMessage(c, mt, msg); err != nil {
			log.Println("Error sending message:", err)
			break
		}
	}
}

func sendMessage(c *websocket.Conn, messageType int, msg []byte) error {
	if err := c.WriteMessage(messageType, msg); err != nil {
		log.Println("Error sending message:", err)
		return err
	}
	return nil
}
