package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	fmt.Println("api project")
	app := fiber.New()

	app.Get("/timestamp", func(c *fiber.Ctx) error {
		timestamp := time.Now().Unix()
		timeString := strconv.Itoa(int(timestamp))
		fmt.Println(timeString)
		return c.JSON(bson.M{"message": timeString})
	})

	app.Listen(":9001")

}
