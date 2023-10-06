package main

import (
	"github.com/labstack/echo/v4"

	"audit/internal/http/handlers"
)

func main() {
	e := echo.New()

	e.GET("/audit", handlers.GetAudit)
	e.POST("/audit", handlers.CreateAudit)

	e.Logger.Fatal(e.Start(":8080"))
}
