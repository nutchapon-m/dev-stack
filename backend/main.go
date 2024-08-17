package main

import (
	"backend-go/api/routes"
	"backend-go/pkgs/configs"
	"backend-go/pkgs/db"
	"backend-go/pkgs/logs"
	"backend-go/pkgs/middlewares"
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	// init configuration
	configs.Init()
	// init database instance
	db.InitDatabase()
	// init log
	logs.Init()
}

func main() {
	app := fiber.New()
	//////////////////////////////////////////////////////////////////////
	// Middleware
	app.Use(middlewares.NewCorsMiddleWare())
	app.Use(middlewares.NewLoggerMiddleWare())

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	})
	//////////////////////////////////////////////////////////////////////
	// Router
	// api routing
	api := app.Group("/api/v1", middlewares.APIKeyAuthMiddleware())
	// user route
	routes.UserRoute(api)

	//////////////////////////////////////////////////////////////////////
	// Gracefully shutting down
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		s := <-c
		if s.String() == "interrupt" {
			fmt.Println("Gracefully shutting down...")
			app.Shutdown()
		}
	}()
	//////////////////////////////////////////////////////////////////////
	// Listen and Serve
	if viper.GetString("app.mode") == "debug" {
		app.Listen("localhost:8000")
	} else {
		app.Listen(":8000")
	}
}
