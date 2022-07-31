package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func emperorLogger(app *echo.Echo) {
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `Time: ${time_rfc3339_nano}` + "\n" +
			`URI: ${uri}` + "\n" +
			`User agent:${user_agent}` + "\n" +
			`Status: ${status}` + "\n" +
			`Error: ${error}` + "\n" +
			`Bytes in: ${bytes_in}, Bytes out:${bytes_out}` + "\n",
	}))
}
