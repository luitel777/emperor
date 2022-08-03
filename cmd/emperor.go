package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
)

// func printHello(c echo.Context) error {
// 	// return c.String(http.StatusOK, "hello world you have "+strconv.Itoa(listFiles())+" files in the data folder")
// 	return c.String(http.StatusOK, "hello world you have "+listFiles()[0]+" files in the data folder")
// }

func startAndExitServer(c *echo.Echo, port int, s *http2.Server) {
	if err := c.StartH2CServer(":"+strconv.Itoa(port), s); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func initializeEcho() {

}

func homePage(app *echo.Echo) {

	renderTemplate(app, "/", "template.html",
		map[string]interface{}{
			"name":         "Aagaman",
			"if":           "ok",
			"fullFilePath": listFiles(),
			"fileName":     listFileNames(),
			"cacheImages":  listCacheFiles(),
			"integer":      []int{1, 2, 3},
		})
}

func checkError(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func main() {

	createFolderIfNotExists()
	createCacheImages()

	app := echo.New()

	// serve from assets directory to /static path
	app.Static("static", "assets")
	app.Static("data", "data")
	app.Static("/", "web")

	port, a := cliemperorFlags()
	if a == true {
		emperorLogger(app)
	}

	s := &http2.Server{
		MaxConcurrentStreams: 500,
		MaxReadFrameSize:     1010100,
		IdleTimeout:          10 * time.Second,
	}
	registerFilepath(app)
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		homePage(app)
		for {
			select {
			case <-ticker.C:
				homePage(app)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	fmt.Printf("\n\n\n")
	listFileNames()

	startAndExitServer(app, port, s)
}
