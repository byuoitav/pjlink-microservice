package controllers

import (
	"net/http"

	"github.com/byuoitav/pjlink-microservice/helpers"
	"github.com/jessemillar/jsonresp"

	"github.com/labstack/echo"
)

func PJLinkRequest(context echo.Context) error {
	request := helpers.PJRequest{}

	err := context.Bind(&request)
	if err != nil {
		return jsonresp.Create(context, http.StatusBadRequest, "Could not read request body: "+err.Error())
	}

	parsedResponse, err := helpers.PJLinkRequest(context.Param("address"),
		context.Param("port"), context.Param("class"), context.Param("passwd"),
		context.Param("command"), context.Param("param"))
	if err != nil {
		// TODO
		return err
	}

	return context.JSON(http.StatusOK, parsedResponse)
}
