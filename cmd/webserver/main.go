package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())  // Ensure this line ends correctly
	e.Use(middleware.Recover()) // Ensure this line ends correctly

	// Routes
	e.GET("/", homeHandler) // Make sure this is correct

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// homeHandler serves the home page
func homeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to My Personal Website!")
}
