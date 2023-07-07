package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"github.com/raymundovr/kvstore/storage"
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

	kvc := c.(*KVContext)
	err := kvc.store.Put(entry.Key, entry.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, KVResponse{Success: false})
	}

	storage.ServiceStorage.WritePut(entry.Key, entry.Value)

	return c.JSON(http.StatusOK, KVResponse{Success: true, Data: *entry})
}

func getHandler(c echo.Context) error {
	key := c.QueryParam("key")

	if key == "" {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	kvc := c.(*KVContext)
	value, err := kvc.store.Get(key)

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
	kvc := c.(*KVContext)
	err := kvc.store.Delete(key)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	storage.ServiceStorage.WriteDelete(key)

	return c.JSON(http.StatusOK, KVResponse{Success: true})

}
