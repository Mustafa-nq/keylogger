package main

import (
	"log"
	"os"
	"path/filepath"

	"keylogger/internal/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Get working directory
	wd, _ := os.Getwd()
	webDir := filepath.Join(wd, "web")
	indexFile := filepath.Join(webDir, "index.html")

	// Validate index.html exists
	if _, err := os.Stat(indexFile); os.IsNotExist(err) {
		log.Fatalf("index.html NOT FOUND at: %s", indexFile)
	}

	// Serve index
	e.File("/", indexFile)

	// Serve /static for CSS & JS
	e.Static("/static", webDir)

	// API routes
	handlers.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
