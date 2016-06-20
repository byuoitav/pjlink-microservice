package handlers

import (
	"net/http"

	"github.com/byuoitav/pjlink-microservice/packages/pjlink"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

func Raw(context echo.Context) error {
	request := pjlink.PJRequest{}

	requestError := context.Bind(&request)
	if requestError != nil {
		return jsonresp.New(context, http.StatusBadRequest, "Could not read request body: "+requestError.Error())
	}

	response, responseError := pjlink.HandleRawRequest(request)
	if responseError != nil {
		return jsonresp.New(context, http.StatusBadRequest, responseError.Error())
	}

	return context.JSON(http.StatusOK, response)
}