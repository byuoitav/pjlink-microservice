package main

import (
	"log"

	"github.com/byuoitav/hateoas"
	"github.com/byuoitav/pjlink-microservice/handlers"
	"github.com/byuoitav/wso2jwt"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := ":8005"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(middleware.CORS())

	err := hateoas.Load("https://raw.githubusercontent.com/byuoitav/pjlink-microservice/master/swagger.json")
	if err != nil {
		log.Fatalln("Could not load swagger.json file. Error: " + err.Error())
	}

	router.Get("/", hateoas.RootResponse)
	router.Get("/health", health.Check)
	router.Get("/raw", handlers.RawInfo, wso2jwt.ValidateJWT())
	router.Post("/raw", handlers.Raw, wso2jwt.ValidateJWT())
	router.Get("/command", handlers.CommandInfo, wso2jwt.ValidateJWT())
	router.Post("/command", handlers.Command, wso2jwt.ValidateJWT())

	log.Println("The PJLink microservice is listening on " + port)
	server := fasthttp.New(port)
	server.ReadBufferSize = 1024 * 10 // Needed to interface properly with WSO2
	router.Run(server)
}
