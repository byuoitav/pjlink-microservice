package controllers

import (
	"net/http"

	"github.com/byuoitav/pjlink-microservice/helpers"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

func Raw(context echo.Context) error {
	request := helpers.PJRequest{}

	requestError := context.Bind(&request)
	if requestError != nil {
		return jsonresp.New(context, http.StatusBadRequest, "Could not read request body: "+requestError.Error())
	}

	response, responseError := helpers.HandleRawRequest(request)
	if responseError != nil {
		return jsonresp.New(context, http.StatusBadRequest, responseError.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func Command(context echo.Context) error {
	request := helpers.PJRequest{}

	requestError := context.Bind(&request)
	if requestError != nil {
		return jsonresp.New(context, http.StatusBadRequest, "Could not read request body: "+requestError.Error())
	}

	response, responseError := helpers.HandleRequest(request)
	if responseError != nil {
		return jsonresp.New(context, http.StatusBadRequest, responseError.Error())
	}

	return context.JSON(http.StatusOK, response)
}
