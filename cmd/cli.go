package main

import (
	"flag"

	"github.com/labstack/echo/v4"
)

func cliInit(app *echo.Echo) int {
	port, a := cliemperorFlags()
	if a == true {
		emperorLogger(app)
	}
	return port

}

func cliemperorFlags() (int, bool) {

	var emperorLogger bool
	flag.BoolVar(&emperorLogger, "log", false, "Enable or disable logging\nfalse (default)\ntrue")

	var emperorPort int
	flag.IntVar(&emperorPort, "port", 8080, "Specify port number")
	flag.Parse()

	return emperorPort, emperorLogger
}

func setPort(portInteger int) int {
	return portInteger
}
