// Manuel Vargas Evans, github.com/github.com/github.com/github.com/mvargasevans@gmail.com

// Package handlers will manage all the REST API requests if needed.
package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//-- GET --//

// Root loads the main application
func Root() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Welcome!")
	}
}

// Fetch will download the artifact
func Fetch() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Fetch!")
	}
}

//-- POST --//

// Create will create an artifact from context
func Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Create")
	}
}
