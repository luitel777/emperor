package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	// this is how map works
	// a := make(map[string]interface{})
	// a["blah"] = "something"
	// log.Println(a["blah"])

	return t.templates.ExecuteTemplate(w, name, data)
}

func registerFilepath(app *echo.Echo) {
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("web/*.html")),
	}
	app.Renderer = renderer
}

func renderTemplate(app *echo.Echo, path string, htmlTmpl string, a any) {
	app.GET(path, func(c echo.Context) error {
		return c.Render(http.StatusOK, htmlTmpl, a)
	})
}
