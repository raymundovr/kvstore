package rest

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/raymundovr/kvstore/core"
)

type KVContext struct {
	echo.Context
	store *core.KVStore
}

func InitializeRest(store *core.KVStore) {
	server := echo.New()
	kvc := &KVContext{store: store}

	server.Use(func (next echo.HandlerFunc) echo.HandlerFunc {
		return func (c echo.Context) error {
			return next(kvc)
		}
	})
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	server.Validator = &KVValidator{validator: validator.New()}

	// Setup routes
	server.PUT("/", putHandler)
	server.GET("/", getHandler)
	server.DELETE("/", deleteHandler)

	server.Logger.Fatal(server.Start(":8080"))
}
