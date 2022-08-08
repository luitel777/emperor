package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
)

func startAndExitServer(c *echo.Echo, port int, s *http2.Server) {
	if err := c.StartH2CServer(":"+strconv.Itoa(port), s); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func homePage(app *echo.Echo) {

	renderTemplate(app, "/", "template.html",
		map[string]interface{}{
			"name":         "Emperor",
			"fullFilePath": listFiles(),
			"fileName":     listFileNames(),
			"cacheImages":  listCacheFiles(),
			"integer":      []int{1, 2, 3},
		})
}

func main() {

	createFolderIfNotExists()
	createCacheImages()

	app := echo.New()

	// serve from assets directory to /static path
	app.Static("static", "assets")
	app.Static("data", "data")
	app.Static("/", "web")

	app.POST("/upload", upload)

	port, a := cliemperorFlags()
	if a == true {
		emperorLogger(app)
	}

	registerFilepath(app)
	homePage(app)

	fmt.Printf("\n\n\n")
	listFileNames()

	app.Start(":" + strconv.Itoa(port))
}
