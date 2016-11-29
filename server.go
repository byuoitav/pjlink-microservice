package main

import (
	"log"
	"net/http"

	"github.com/byuoitav/authmiddleware"
	"github.com/byuoitav/hateoas"
	"github.com/byuoitav/pjlink-microservice/handlers"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := hateoas.Load("https://raw.githubusercontent.com/byuoitav/pjlink-microservice/master/swagger.json")
	if err != nil {
		log.Fatalln("Could not load swagger.json file. Error: " + err.Error())
	}

	port := ":8005"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(middleware.CORS())

	// Use the `secure` routing group to require authentication
	secure := router.Group("", echo.WrapMiddleware(authmiddleware.Authenticate))

	router.GET("/", echo.WrapHandler(http.HandlerFunc(hateoas.RootResponse)))
	router.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))

	secure.GET("/raw", handlers.RawInfo)
	secure.POST("/raw", handlers.Raw)
	secure.GET("/command", handlers.CommandInfo)
	secure.POST("/command", handlers.Command)

	secure.GET("/:address/power/on", handlers.PowerOn)
	secure.GET("/:address/power/standby", handlers.PowerOff)
	secure.GET("/:address/display/blank", handlers.DisplayBlank)
	secure.GET("/:address/display/unblank", handlers.DisplayUnBlank)
	secure.GET("/:address/volume/mute", handlers.VolumeMute)
	secure.GET("/:address/volume/unmute", handlers.VolumeUnMute)
	secure.GET("/:address/input/:port", handlers.SetInputPort)

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
