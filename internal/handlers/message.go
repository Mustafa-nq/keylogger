package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MessageReq struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

func HandleMessage(c echo.Context) error {
	var req MessageReq
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid JSON")
	}

	fmt.Printf("[CHAT] %s (%s): %s\n", req.User, c.RealIP(), req.Message)

	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
