package handlers

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo) {
	e.POST("/api/message", HandleMessage)
	e.POST("/api/keystrokes", HandleKeystrokes)
}
