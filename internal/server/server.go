// Manuel Vargas Evans, mvargasevans@gmail.com

// Package server handles the configuration of the Echo server.
//
// To configure the certificates, these have been obtained from Let's Encrypt
// as self-signed. You can obtain your own or from a proper CA if required.
package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Configuration variables for the server
const (
	HTTPPort  = ":80"
	HTTPSPort = ":443"
	CERT      = "../../internal/server/localhost.crt"
	KEY       = "../../internal/server/localhost.key"
)

// RedirectHTTP starts a parallel server that monitors port 80 and redirects TLS
func redirectHTTP() {
	e := echo.New()
	e.Use(middleware.HTTPSRedirect())
	go func() { e.Logger.Fatal(e.Start(HTTPPort)) }()
}

// Start will start the server
func Start(e *echo.Echo, TLS bool) {
	if TLS {
		redirectHTTP()
		e.Logger.Fatal(e.StartTLS(HTTPSPort, CERT, KEY))
	} else {
		e.Logger.Fatal(e.Start(HTTPPort))
	}
}
