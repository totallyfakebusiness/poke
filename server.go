package main

import (
	_ "embed"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/smira/go-statsd"
)

//go:embed poke.html
var index string

func main() {
	// init statsd
	statsHost, ok := os.LookupEnv("STATSD_HOST")
	if !ok {
		statsHost = "localhost"
	}
	statsClient := statsd.NewClient(statsHost + ":8125")
	defer statsClient.Close()

	// init echo
	e := echo.New()
	e.Use(middleware.Logger())

	// index page
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, index)
	})

	// basic health endpoint
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "👍")
	})

	// emit a metric when poked
	e.POST("/poke", func(c echo.Context) error {
		statsClient.Incr("pokes", 1)
		return c.NoContent(http.StatusNoContent)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
