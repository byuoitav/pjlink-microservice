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
		jsonresp.New(context.Response(), http.StatusBadRequest, "Could not read request body: "+requestError.Error())
		return nil
	}

	response, responseError := pjlink.HandleRequest(request)
	if responseError != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, responseError.Error())
		return nil
	}

	if (response.Response[0] == "unavailable time") || (response.Response[0] == "device failure") {
		jsonresp.New(context.Response(), http.StatusInternalServerError, response.Response[0])
		return nil
	}

	return context.JSON(http.StatusOK, response)
}

func CommandInfo(context echo.Context) error {
	jsonresp.New(context.Response(), http.StatusBadRequest, "Send a POST request to the /command endpoint with a body including Address, Port, Class, Password, Command, and Parameter tokens")
	return nil
}
