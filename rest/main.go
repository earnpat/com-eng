package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

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

	////////// mock data //////////
	jsonFile, err := os.Open("../todo.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonData, _ := io.ReadAll(jsonFile)

	var todoData []TodoData
	err = json.Unmarshal(jsonData, &todoData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//////////////////////////////

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
			"todo":      todoData,
		})
	})

	app.Listen(":9001")
}
