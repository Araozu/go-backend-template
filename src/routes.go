package src

import (
	"embed"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Sets up the Echo server, and registers all routes and sub routes
func (s *Server) RegisterRoutes(staticEmbed *embed.FS) http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(middleware.CSRF())

	// Static files
	fileServer := http.FileServer(http.FS(staticEmbed))
	e.GET("/public/*", echo.WrapHandler(fileServer))

	// TODO: Register subroutes here
	// login.SetupRoutes(e.Group("/auth"))

	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello")
	})

	return e
}
