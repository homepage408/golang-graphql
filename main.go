package main

import (
	"fmt"
	"log"
	"strconv"

	cf "github.com/golang-graphql/config"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var (
	err  error
	d    struct{}
	conf = cf.Config{}
	port = "8080"
)

func run() error {
	if err = godotenv.Load(); err != nil {
		return fmt.Errorf("failed to load .env")
	}

	conf.Init()
	// ctx := context.Background()
	// ctx, cancel := context.WithCancel(ctx)
	// defer cancel()

	server := fiber.New(
		fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"Status":    fiber.StatusInternalServerError,
					"IsSuccess": false,
					"Message":   err.Error(),
					"Data":      d,
				})
			},
			AppName: conf.Service.Name,
		},
	)

	server.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"Status":    fiber.StatusOK,
			"IsSuccess": true,
			"Message":   conf.Service.Name,
			"Data":      d,
		})
	})

	if conf.Service.Port != 0 {
		port = strconv.Itoa(conf.Service.Port)
	}

	return server.Listen(":" + port)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
