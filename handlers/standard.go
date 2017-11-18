package handlers

import (
	"net/http"
	"os"
	"strings"

	"github.com/byuoitav/av-api/statusevaluators"
	"github.com/byuoitav/pjlink-microservice/pjlink"
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

	if contains("success", response.Response) || contains("OK", response.Response) {
		return context.JSON(http.StatusOK, statusevaluators.PowerStatus{"on"})
	}

	return context.JSON(http.StatusInternalServerError, response)
}

func PowerOff(context echo.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "power", "power-off")

	response, responseError := pjlink.HandleRequest(request)
	if responseError != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, responseError.Error())
		return nil
	}

	if contains("success", response.Response) || contains("OK", response.Response) {
		return context.JSON(http.StatusOK, statusevaluators.PowerStatus{"standby"})
	}

	return context.JSON(http.StatusOK, response)
}

func GetPowerStatus(context echo.Context) error {

	request := formRequestFromEnvVars(context.Param("address"), "power", "query")

	response, err := pjlink.GetPowerStatus(request)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
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

	if contains("success", response.Response) || contains("OK", response.Response) {
		return context.JSON(http.StatusOK, statusevaluators.BlankedStatus{true})
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

	if contains("success", response.Response) || contains("OK", response.Response) {
		return context.JSON(http.StatusOK, statusevaluators.BlankedStatus{false})
	}

	return context.JSON(http.StatusOK, response)
}

func GetBlankedStatus(context echo.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "av-mute", "query")

	response, err := pjlink.GetBlankedStatus(request)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
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

	if contains("success", response.Response) || contains("OK", response.Response) {
		return context.JSON(http.StatusOK, statusevaluators.MuteStatus{true})
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

	if contains("success", response.Response) || contains("OK", response.Response) {
		return context.JSON(http.StatusOK, statusevaluators.MuteStatus{false})
	}

	return context.JSON(http.StatusOK, response)
}

func GetMuteStatus(context echo.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "av-mute", "query")

	response, err := pjlink.GetMuteStatus(request)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
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

	if contains("success", response.Response) || contains("OK", response.Response) {
		return context.JSON(http.StatusOK, statusevaluators.Input{context.Param("port")})
	}

	return context.JSON(http.StatusInternalServerError, response)
}

func GetCurrentInput(context echo.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "input", "query")

	response, err := pjlink.GetCurrentInput(request)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func GetInputList(context echo.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "input-list", "query")

	response, err := pjlink.GetInputList(request)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
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

func contains(str string, list []string) bool {

	for _, s := range list {

		if strings.Contains(s, str) {
			return true
		}
	}

	return false
}
