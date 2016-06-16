package main

import (
	"fmt"

	"github.com/byuoitav/pjlink-microservice/controllers"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := ":8005"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())

	/*
		err := hateoas.Load("https://raw.githubusercontent.com/byuoitav/pjlink-microservice/master/swagger.yml")
		if err != nil {
			fmt.Printf("Could not load swagger.yaml file. Error: %s", err.Error())
			panic(err)
		}
	*/

	//router.Get("/", hateoas.RootResponse)

	router.Get("/health", health.Check)

	router.Post("/raw", controllers.Raw)

	fmt.Printf("The PJLink microservice is listening on %s\n", port)
	server := fasthttp.New(port)
	server.ReadBufferSize = 1024 * 10 // Needed to interface properly with WSO2
	router.Run(server)
}
