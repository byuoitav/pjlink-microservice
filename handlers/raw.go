package handlers

import (
	"net/http"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/v2/auth"
	"github.com/byuoitav/pjlink-microservice/pjlink"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

func Raw(context echo.Context) error {
	if ok, err := auth.CheckAuthForLocalEndpoints(context, "write-state"); !ok {
		if err != nil {
			log.L.Warnf("Problem getting auth: %v", err.Error())
		}
		return context.String(http.StatusUnauthorized, "unauthorized")
	}

	request := pjlink.PJRequest{}

	requestError := context.Bind(&request)
	if requestError != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, "Could not read request body: "+requestError.Error())
		return nil
	}

	response, responseError := pjlink.HandleRawRequest(request)
	if responseError != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, responseError.Error())
		return nil
	}

	return context.JSON(http.StatusOK, response)
}

func RawInfo(context echo.Context) error {
	if ok, err := auth.CheckAuthForLocalEndpoints(context, "write-state"); !ok {
		if err != nil {
			log.L.Warnf("Problem getting auth: %v", err.Error())
		}
		return context.String(http.StatusUnauthorized, "unauthorized")
	}

	jsonresp.New(context.Response(), http.StatusBadRequest, "Send a POST request to the /raw endpoint with a body including Address, Port, Class, Password, Command, and Parameter tokens")
	return nil
}
