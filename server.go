// pjlinkService project main.go
package main

import (
	"fmt"

	//"github.com/byuoitav/hateoas"
	"github.com/byuoitav/pjlink-service/controllers"

	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
)

func main() {

	port := ":8005"
	echoServer := echo.New()
	echoServer.Pre(middleware.RemoveTrailingSlash())

	/*
		err := hateoas.Load("https://raw.githubusercontent.com/byuoitav/pjlink-service/master/swagger.yml")
		if err != nil {
			fmt.Printf("Could not load swagger.yaml file. Error: %s", err.Error())
			panic(err)
		}
	*/

	//echoServer.Get("/", hateoas.RootResponse)

	echoServer.Get("/health", health.Check)

	echoServer.Get("/address/:address/port/:port/class/:class/passwd/:passwd/"+
		"command/:command/param/:param", controllers.PjlinkRequest)

	echoServer.Run(fasthttp.New(port))
}

func handleResponse(response string) {
	//TODO do something with response
	fmt.Println(response)
}
