package main

import (
	"log"

	"github.com/byuoitav/pjlink-microservice/handlers"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := ":8005"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())

	// err := hateoas.Load("https://raw.githubusercontent.com/byuoitav/pjlink-microservice/master/swagger.json")
	// if err != nil {
	// 	log.Fatalln("Could not load swagger.json file. Error: " + err.Error())
	// }

	//router.Get("/", hateoas.RootResponse)

	router.Get("/health", health.Check)

	router.Post("/raw", handlers.Raw)

	router.Post("/command", handlers.Command)

	log.Println("The PJLink microservice is listening on " + port)
	router.Run(fasthttp.New(port))
}
