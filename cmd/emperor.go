package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
)

func printHello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world you have "+strconv.Itoa(listFiles())+" files in the data folder")
}

func startAndExitServer(c *echo.Echo, port int, s *http2.Server) {
	if err := c.StartH2CServer(":"+strconv.Itoa(port), s); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func initializeEcho() {

}

func main() {

	createFolderIfNotExists()

	app := echo.New()
	var port int = cliInit(app)

	s := &http2.Server{
		MaxConcurrentStreams: 500,
		MaxReadFrameSize:     1010100,
		IdleTimeout:          10 * time.Second,
	}

	app.GET("/", printHello)

	startAndExitServer(app, port, s)
}
