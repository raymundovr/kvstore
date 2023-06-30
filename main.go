package main

import (
	"fmt"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize events storage before server
	if err := initializeEventsStorage(); err != nil {
		fmt.Println("Cannot start service: %w", err)
		// At the moment we want to crash
		os.Exit(1)
	}

	server := echo.New()

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	server.Validator = &KVValidator{validator: validator.New()}

	// Setup routes
	server.PUT("/", putHandler)
	server.GET("/", getHandler)
	server.DELETE("/", deleteHandler)

	server.Logger.Fatal(server.Start(":12345"))
}
