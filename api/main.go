package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("api project")
	app := fiber.New()

	app.Get("/timestamp", func(c *fiber.Ctx) error {
		timestamp := time.Now().Unix()
		timeString := strconv.Itoa(int(timestamp))
		fmt.Println(timeString)
		return c.SendString(timeString)
	})

	app.Listen(":9000")

}
