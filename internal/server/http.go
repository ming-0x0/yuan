package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
}

func NewServer() *Server {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	return &Server{
		echo: e,
	}
}

func (s *Server) Start(addr string) error {
	return s.echo.Start(addr)
}

func (s *Server) Echo() *echo.Echo {
	return s.echo
}
