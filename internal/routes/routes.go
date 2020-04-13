// Manuel Vargas Evans, github.com/mvargasevans@gmail.com

// Package routes sets the routes for the application
package routes

import (
	"github.com/mvargasevans/goecho-setup/internal/handlers"

	"github.com/labstack/echo/v4"
)

// Set will start the REST API routes.
func Set(e *echo.Echo) {

	//-- GET --//
	e.GET("/", handlers.Root())
	e.GET("/fetch", handlers.Fetch())

	//-- POST --//
	e.POST("/create", handlers.Create())
	return
}
