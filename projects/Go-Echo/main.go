package main

import (
	"github.com/Devisrisamidurai/Go-echo/cmd/api/handlers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/health-check", handlers.HealthCheckHandler)
	e.GET("/posts", handlers.PostIndexHandler)
	e.GET("/post/:id", handlers.PostSingleHandler)

	e.Logger.Fatal(e.Start(":1323"))

}
