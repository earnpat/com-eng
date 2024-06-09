package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	fmt.Println("rest start")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		var query struct {
			Timestamp string `query:"timestamp" json:"timestamp"`
		}
		if err := c.QueryParser(&query); err != nil {
			return err
		}

		timestampInt, timestampIntErr := strconv.ParseInt(query.Timestamp, 10, 64)
		if timestampIntErr != nil {
			return timestampIntErr
		}

		return c.Status(fiber.StatusOK).JSON(bson.M{
			"timestamp": timestampInt,
		})
	})

	app.Listen(":9001")
}
