package main

import (
	"github.com/labstack/echo"
	"github.com/y-tkbridge/youtube-manager-go/web/api"
)

func main() {
	e := echo.New()
	routes.Inite(e)

	e.Logger.Fatal(e.Start(":8080"))
}