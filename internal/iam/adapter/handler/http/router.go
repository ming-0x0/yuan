package http

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(g *echo.Group, authHandler *AuthHandler) {
	authGroup := g.Group("/auth")
	authGroup.POST("/register", authHandler.Register)
	authGroup.POST("/login", authHandler.Login)
}

func RegisterUserRoutes(g *echo.Group, userHandler *UserHandler) {
	userGroup := g.Group("/users")
	// TODO: Move secret to config
	userGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	}))
	userGroup.GET("/me", userHandler.GetProfile)
	userGroup.PUT("/me", userHandler.UpdateProfile)
}
