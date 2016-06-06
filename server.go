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

	router.Post("/command", controllers.Command)

	fmt.Printf("The PJLink microservice is listening on %s\n", port)
	router.Run(fasthttp.New(port))
}
