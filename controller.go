package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type KVEntry struct {
	Key   string `json:"key" query:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type KVResponse struct {
	Success bool    `json:"success"`
	Data    KVEntry `json:"data"`
}

func putHandler(c echo.Context) error {
	entry := new(KVEntry)

	if err := c.Bind(entry); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, KVResponse{Success: false})
	}

	if err := c.Validate(entry); err != nil {
		return err
	}

	err := Put(entry.Key, entry.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, KVResponse{Success: false})
	}

	return c.JSON(http.StatusOK, KVResponse{Success: true, Data: *entry})
}

func getHandler(c echo.Context) error {
	c.Logger().Print("Arriving")
	key := c.QueryParam("key")

	if key == "" {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	value, err := Get(key)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	entry := KVEntry{Key: key, Value: value}

	return c.JSON(http.StatusOK, KVResponse{Success: true, Data: entry})

}

func deleteHandler(c echo.Context) error {
	c.Logger().Print("Arriving")
	key := c.QueryParam("key")

	if key == "" {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err := Delete(key)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, KVResponse{Success: true})

}
