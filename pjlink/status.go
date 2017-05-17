package pjlink

import (
	"strings"

	"github.com/byuoitav/av-api/status"
)

func GetPowerStatus(request PJRequest) (status.PowerStatus, error) {

	response, err := HandleRequest(request)
	if err != nil {
		return status.PowerStatus{}, err
	}

	var output status.PowerStatus
	if strings.EqualFold(response.Response[0], "power-on (lamp on)") {
		output.Power = "on"
	} else {
		output.Power = "standby"
	}
	return output, nil
}

func GetBlankedStatus(request PJRequest) (status.BlankedStatus, error) {

	response, err := HandleRequest(request)
	if err != nil {
		return status.BlankedStatus{}, err
	}

	var output status.BlankedStatus
	if strings.EqualFold(response.Response[0], "video and audio mute on") ||
		strings.EqualFold(response.Response[0], "video mute on, audio mute off") {
		output.Blanked = true
	}

	return output, nil
}

func GetMuteStatus(request PJRequest) (status.MuteStatus, error) {

	response, err := HandleRequest(request)
	if err != nil {
		return status.MuteStatus{}, err
	}

	var output status.MuteStatus
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

func GetInputList(request PJRequest) (status.VideoList, error) {

	response, err := HandleRequest(request)
	if err != nil {
		return status.VideoList{}, err
	}

	var output status.VideoList
	for _, entry := range response.Response {

		output.Inputs = append(output.Inputs, status.Input{Input: entry})

	}

	return output, nil
}
