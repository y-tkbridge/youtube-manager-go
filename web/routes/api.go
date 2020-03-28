package routes

import (
	"github.com/labstack/echo"
<<<<<<< HEAD
	"github.com/y-tkbridge/youtube-manager-go/web/api"
=======
	"github.com/y-bridge/youtube-manager-go/web/api"
>>>>>>> 修正内容の反映
)

func Init(e *echo.Echo) {
	g := e.Group("/api")
	{
		g.GET("/popular", api.GetchMostPopularVideos())
	}
}
