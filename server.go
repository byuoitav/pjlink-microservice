package main

import (
	"github.com/byuoitav/common"
	"net/http"

	"github.com/byuoitav/hateoas"
	"github.com/byuoitav/pjlink-microservice/handlers"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := ":8005"
	router := common.NewRouter()

	// Use the `secure` routing group to require authentication
	write := router.Group("", auth.AuthorizeRequest("write-state", "room", auth.LookupResourceFromAddress))
	read := router.Group("", auth.AuthorizeRequest("read-state", "room", auth.LookupResourceFromAddress)

	router.GET("/", echo.WrapHandler(http.HandlerFunc(hateoas.RootResponse)))
	router.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))
	router.GET("/status", echo.WrapHandler(http.HandlerFunc(health.Check)))

	read.GET("/raw", handlers.RawInfo)
	read.POST("/raw", handlers.Raw)
	read.GET("/command", handlers.CommandInfo)
	read.POST("/command", handlers.Command)

	//status endpoints
	read.GET("/:address/power/status", handlers.GetPowerStatus)
	read.GET("/:address/display/status", handlers.GetBlankedStatus)
	read.GET("/:address/volume/mute/status", handlers.GetMuteStatus)
	read.GET("/:address/input/current", handlers.GetCurrentInput)
	read.GET("/:address/input/list", handlers.GetInputList)

	//functionality endpoints
	write.GET("/:address/power/on", handlers.PowerOn)
	write.GET("/:address/power/standby", handlers.PowerOff)
	write.GET("/:address/display/blank", handlers.DisplayBlank)
	write.GET("/:address/display/unblank", handlers.DisplayUnBlank)
	write.GET("/:address/volume/mute", handlers.VolumeMute)
	write.GET("/:address/volume/unmute", handlers.VolumeUnMute)
	write.GET("/:address/input/:port", handlers.SetInputPort)

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
