package handlers

import (
	"net/http"

	"github.com/byuoitav/pjlink-microservice/packages/pjlink"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

func Command(context echo.Context) error {
	request := pjlink.PJRequest{}

	requestError := context.Bind(&request)
	if requestError != nil {
		return jsonresp.New(context, http.StatusBadRequest, "Could not read request body: "+requestError.Error())
	}

	response, responseError := pjlink.HandleRequest(request)
	if responseError != nil {
		return jsonresp.New(context, http.StatusBadRequest, responseError.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func CommandInfo(context echo.Context) error {
	return jsonresp.New(context, http.StatusBadRequest, "Send a POST request to the /command endpoint with a body including Address, Port, Class, Password, Command, and Parameter tokens")
}

func PowerOn(context echo.Context) error {
	address := context.Param("address")

	request := pjlink.PJRequest{Address: address, Port: "4352", Password: "test", Class: "1", Command: "power", Parameter: "power-on"}
	response, responseError := pjlink.HandleRequest(request)
	if responseError != nil {
		return jsonresp.New(context, http.StatusBadRequest, responseError.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func PowerOff(context echo.Context) error {
	address := context.Param("address")
	//TODO: Pull projector password out into a variable.
	request := pjlink.PJRequest{Address: address, Port: "4352", Password: "test", Class: "1", Command: "power", Parameter: "power-off"}
	response, responseError := pjlink.HandleRequest(request)
	if responseError != nil {
		return jsonresp.New(context, http.StatusBadRequest, responseError.Error())
	}

	return context.JSON(http.StatusOK, response)
}
