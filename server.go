package main

import (
	"github.com/amaru0601/go-rent/db"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	entClient, err := db.GetClient()
	logger := e.Logger
	logger.Print("GO")
	if err != nil {
		logger.Info("Database could not initialize")
	}

	defer entClient.Close()
}
