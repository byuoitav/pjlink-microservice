package controllers

import (
	"net/http"

	"github.com/byuoitav/pjlink-microservice/helpers"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

func Raw(context echo.Context) error {
	request := helpers.RawPJRequest{}

	err := context.Bind(&request)
	if err != nil {
		return jsonresp.New(context, http.StatusBadRequest, "Could not read request body: "+err.Error())
	}

	response, err := helpers.HandleRawRequest(request)
	if err != nil {
		return jsonresp.New(context, http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
