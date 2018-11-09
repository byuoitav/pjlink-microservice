package main

import (
	"net/http"

	"github.com/byuoitav/hateoas"
	"github.com/byuoitav/pjlink-microservice/handlers"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := ":8005"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(middleware.CORS())

	// Use the `secure` routing group to require authentication
	//secure := router.Group("", echo.WrapMiddleware(authmiddleware.Authenticate))

	router.GET("/", echo.WrapHandler(http.HandlerFunc(hateoas.RootResponse)))
	router.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))
	router.GET("/status", echo.WrapHandler(http.HandlerFunc(health.Check)))

	router.GET("/raw", handlers.RawInfo)
	router.POST("/raw", handlers.Raw)
	router.GET("/command", handlers.CommandInfo)
	router.POST("/command", handlers.Command)

	//status endpoints
	router.GET("/:address/power/status", handlers.GetPowerStatus)
	router.GET("/:address/display/status", handlers.GetBlankedStatus)
	router.GET("/:address/volume/mute/status", handlers.GetMuteStatus)
	router.GET("/:address/input/current", handlers.GetCurrentInput)
	router.GET("/:address/input/list", handlers.GetInputList)

	//functionality endpoints
	router.GET("/:address/power/on", handlers.PowerOn)
	router.GET("/:address/power/standby", handlers.PowerOff)
	router.GET("/:address/display/blank", handlers.DisplayBlank)
	router.GET("/:address/display/unblank", handlers.DisplayUnBlank)
	router.GET("/:address/volume/mute", handlers.VolumeMute)
	router.GET("/:address/volume/unmute", handlers.VolumeUnMute)
	router.GET("/:address/input/:port", handlers.SetInputPort)

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
