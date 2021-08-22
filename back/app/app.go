package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() {
	print("started")
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET(`/`, hello)
	e.Logger.Fatal(e.Start(":3000"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world")
}
