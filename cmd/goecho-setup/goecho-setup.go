// Manuel Vargas Evans, mvargasevans@gmail.com

// Package goecho-setup is a webserver based on Echo Framework.
// https://echo.labstack.com/.

// goecho-setup can be used to manufacture artifacts for the goecho.

package main

import (
	"mvargasevans/goecho-setup/internal/routes"
	"mvargasevans/goecho-setup/internal/server"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Create a new instance of Echo
	e := echo.New()

	//-- Middlewares --//
	// Basic middleware for logging and recovering over failures
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Configuration middleware to enable security features
	e.Use(middleware.Secure())

	//-- Routes --//
	routes.Set(e)

	//-- Server --//
	server.Start(e, true)
}
