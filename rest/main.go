package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type TodoData struct {
	Id        int64  `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
	UserId    int64  `json:"userId"`
}

func main() {
	fmt.Println("rest start")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(bson.M{"ok": true})
	})

	app.Get("/response", func(c *fiber.Ctx) error {
		jsonFile, err := os.Open("../todo.json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		jsonData, _ := io.ReadAll(jsonFile)

		var todoData []TodoData
		err = json.Unmarshal(jsonData, &todoData)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(bson.M{
			"todo": todoData,
		})
	})

	app.Listen(":9001")
}
