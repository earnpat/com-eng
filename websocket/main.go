package main

import (
	"fmt"
	"io"
	"log"
	"os"

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
			_, _, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("read:", err)
				}
				break
			}

			if err = c.WriteMessage(websocket.TextMessage, nil); err != nil {
				log.Println("write:", err)
			}
		}
	}))

	app.Get("/ws/response", websocket.New(func(c *websocket.Conn) {
		for {
			_, _, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("read:", err)
				}
				break
			}

			jsonFile, err := os.Open("../todo.json")
			if err != nil {
				fmt.Println(err)
			}
			defer jsonFile.Close()

			jsonData, _ := io.ReadAll(jsonFile)

			if err = c.WriteMessage(websocket.BinaryMessage, jsonData); err != nil {
				log.Println("write:", err)
			}
		}
	}))

	log.Fatal(app.Listen(":9003"))
}
