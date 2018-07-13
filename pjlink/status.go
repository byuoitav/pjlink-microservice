package pjlink

import (
	"strings"

	se "github.com/byuoitav/av-api/statusevaluators"
	"github.com/byuoitav/common/status"
)

func GetPowerStatus(request PJRequest) (status.Power, error) {

	response, err := HandleRequest(request)
	if err != nil {
		return status.Power{}, err
	}

	var output status.Power
	if strings.EqualFold(response.Response[0], "power-on (lamp on)") {
		output.Power = "on"
	} else {
		output.Power = "standby"
	}
	return output, nil
}

func GetBlankedStatus(request PJRequest) (status.Blanked, error) {

	response, err := HandleRequest(request)
	if err != nil {
		return status.Blanked{}, err
	}

	var output status.Blanked
	if strings.EqualFold(response.Response[0], "video and audio mute on") ||
		strings.EqualFold(response.Response[0], "video mute on, audio mute off") {
		output.Blanked = true
	}

	return output, nil
}

func GetMuteStatus(request PJRequest) (status.Mute, error) {

	response, err := HandleRequest(request)
	if err != nil {
		return status.Mute{}, err
	}

	var output status.Mute
	if strings.EqualFold(response.Response[0], "video and audio mute on") ||
		strings.EqualFold(response.Response[0], "audio mute on, video mute off") {
		output.Muted = true
	}

	return output, nil
}

func GetCurrentInput(request PJRequest) (status.Input, error) {

	response, err := HandleRequest(request)
	if err != nil {
		return status.Input{}, err
	}

	return status.Input{Input: response.Response[0]}, nil
}

func GetInputList(request PJRequest) (se.VideoList, error) {

	response, err := HandleRequest(request)
	if err != nil {
		return se.VideoList{}, err
	}

	var output se.VideoList
	for _, entry := range response.Response {

		output.Inputs = append(output.Inputs, status.Input{Input: entry})

	}

	return output, nil
}
