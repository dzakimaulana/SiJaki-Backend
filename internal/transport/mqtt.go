package transport

// https://www.emqx.com/en/blog/how-to-use-mqtt-in-golang

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gofiber/websocket/v2"
)

func Subscribe(client mqtt.Client, c *websocket.Conn) {
	topic := "sijaki/fullness"

	client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("Received MQTT message: %s", msg.Payload())
		sendMessage(c, websocket.TextMessage, []byte(msg.Payload()))
	})

	log.Printf("✅ Successfully subscribed to topic: %s\n", topic)
}
