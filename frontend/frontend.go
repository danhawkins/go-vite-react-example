package frontend

import (
	"embed"

	"github.com/labstack/echo/v4"
)

var (
	//go:embed dist/*
	dist embed.FS

	//go:embed dist/index.html
	indexHTML embed.FS

	distDirFS     = echo.MustSubFS(dist, "dist")
	distIndexHTML = echo.MustSubFS(indexHTML, "dist")
)

func RegisterHandlers(e *echo.Echo) {
	// Use the static assets from the dist directory
	e.FileFS("/", "index.html", distIndexHTML)
	e.StaticFS("/", distDirFS)
}
