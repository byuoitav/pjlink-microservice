package controllers

import (
	"net/http"

	"github.com/byuoitav/pjlink-microservice/helpers"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

func Command(context echo.Context) error {
	request := helpers.PJRequest{}

	err := context.Bind(&request)
	if err != nil {
		return jsonresp.Create(context, http.StatusBadRequest, "Could not read request body: "+err.Error())
	}

	response, err := helpers.HandleRequest(request)
	if err != nil {
		return jsonresp.Create(context, http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
