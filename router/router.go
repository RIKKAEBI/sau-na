package router

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func InitRouter(e *echo.Echo) {
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("template/*.html")),
	}
	e.Renderer = renderer

	e.File("/", "./components/dist/index.html")
	e.File("/favicon.ico", "./components/dist/favicon.ico")
	e.File("/assets/main.js", "./components/dist/assets/main.js")
	e.File("/assets/style.css", "./components/dist/assets/style.css")
}
