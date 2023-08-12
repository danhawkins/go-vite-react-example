package main

import (
	"fmt"
	"net/http"

	"github.com/danhawkins/go-vite-react-example/frontend"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create a new echo server
	e := echo.New()

	// Add standard middleware
	e.Use(middleware.Logger())

	// Setup the frontend handlers to service vite static assets
	frontend.RegisterHandlers(e)

	// Setup the API Group
	api := e.Group("/api")

	// Basic APi endpoint
	api.GET("/message", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, from the golang World!"})
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", 3000)))
}
