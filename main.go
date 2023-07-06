package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	cf "github.com/golang-graphql/config"
	"github.com/golang-graphql/graph"
	"github.com/golang-graphql/graph/resolver"
	"github.com/valyala/fasthttp/fasthttpadaptor"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var (
	err  error
	d    struct{}
	conf = cf.Config{}
	port = "8089"
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

	// server.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 		"Status":    fiber.StatusOK,
	// 		"IsSuccess": true,
	// 		"Message":   conf.Service.Name,
	// 		"Data":      d,
	// 	})
	// })

	handlerGraphql(server)

	if conf.Service.Port != 0 {
		port = strconv.Itoa(conf.Service.Port)
	}

	return server.Listen(":" + port)
}

func handlerGraphql(srv *fiber.App) {
	server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &resolver.Resolver{},
	}))

	server.AddTransport(transport.Websocket{})
	server.AddTransport(transport.GET{})
	server.AddTransport(transport.POST{})

	fiberHandler := fasthttpadaptor.NewFastHTTPHandler(server)

	srv.All("/", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(playground.Handler("GraphQL playground", "/query"))(c.Context())
		return nil
	})

	srv.All("/query", func(c *fiber.Ctx) error {
		fiberHandler(c.Context())
		return nil
	})
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
