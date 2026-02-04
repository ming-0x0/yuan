package http

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(g *echo.Group, handler *BlogHandler) {
	blogGroup := g.Group("/blogs")

	// Public routes
	blogGroup.GET("/:id", handler.Get)
	blogGroup.GET("", handler.List)

	// Protected routes
	// TODO: Move secret to config
	blogGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	}))
	blogGroup.POST("", handler.Create)
	blogGroup.PUT("/:id", handler.Update)
	blogGroup.DELETE("/:id", handler.Delete)
}
