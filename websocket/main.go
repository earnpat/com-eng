package main

import (
	"fmt"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("websocket start")
	app := fiber.New()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/test", websocket.New(func(c *websocket.Conn) {
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
			}

			if err = c.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println("write:", err)
			}
		}
	}))

	log.Fatal(app.Listen(":9003"))
}
