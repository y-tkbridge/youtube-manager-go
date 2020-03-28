package main

import (
	"github.com/labstack/echo"
	"github.com/y-tkbridge/youtube-manager-go/web/routes"
)

func main() {
	e := echo.New()
	routes.Init(e)

	e.Logger.Fatal(e.Start(":8080"))
}
