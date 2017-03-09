package handlers

import (
	"net/http"
	"os"

	"github.com/byuoitav/pjlink-microservice/packages/pjlink"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

func PowerOn(context echo.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "power", "power-on")

	response, responseError := pjlink.HandleRequest(request)
	if responseError != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, responseError.Error())
		return nil
	}

	return context.JSON(http.StatusOK, response)
}

func PowerOff(context echo.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "power", "power-off")

	response, responseError := pjlink.HandleRequest(request)
	if responseError != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, responseError.Error())
		return nil
	}

	return context.JSON(http.StatusOK, response)
}

//some projectors *panasonic - cough* only accept av mute, not just blank, so
//a blank command both blanks and mutes
func DisplayBlank(context echo.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "av-mute", "av-mute-on")

	response, responseError := pjlink.HandleRequest(request)
	if responseError != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, responseError.Error())
		return nil
	}

	return context.JSON(http.StatusOK, response)
}

func DisplayUnBlank(context echo.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "av-mute", "av-mute-off")

	response, responseError := pjlink.HandleRequest(request)
	if responseError != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, responseError.Error())
		return nil
	}

	return context.JSON(http.StatusOK, response)
}

func VolumeMute(context echo.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "av-mute", "audio-mute-on")

	response, responseError := pjlink.HandleRequest(request)
	if responseError != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, responseError.Error())
		return nil
	}

	return context.JSON(http.StatusOK, response)
}

func VolumeUnMute(context echo.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "av-mute", "audio-mute-off")

	response, responseError := pjlink.HandleRequest(request)
	if responseError != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, responseError.Error())
		return nil
	}

	return context.JSON(http.StatusOK, response)
}

func SetInputPort(context echo.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "input", context.Param("port"))

	response, responseError := pjlink.HandleRequest(request)
	if responseError != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, responseError.Error())
		return nil
	}

	return context.JSON(http.StatusOK, response)
}

func formRequestFromEnvVars(address, command, parameter string) pjlink.PJRequest {
	request := pjlink.PJRequest{
		Address:   address,
		Port:      os.Getenv("PJLINK_PORT"),
		Password:  os.Getenv("PJLINK_PASS"),
		Class:     "1",
		Command:   command,
		Parameter: parameter,
	}

	return request
}
