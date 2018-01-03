package pjlink

import (
	"strings"

	se "github.com/byuoitav/av-api/statusevaluators"
)

func GetPowerStatus(request PJRequest) (se.PowerStatus, error) {

	response, err := HandleRequest(request)
	if err != nil {
		return se.PowerStatus{}, err
	}

	var output se.PowerStatus
	if strings.EqualFold(response.Response[0], "power-on (lamp on)") {
		output.Power = "on"
	} else {
		output.Power = "standby"
	}
	return output, nil
}

func GetBlankedStatus(request PJRequest) (se.BlankedStatus, error) {

	response, err := HandleRequest(request)
	if err != nil {
		return se.BlankedStatus{}, err
	}

	var output se.BlankedStatus
	if strings.EqualFold(response.Response[0], "video and audio mute on") ||
		strings.EqualFold(response.Response[0], "video mute on, audio mute off") {
		output.Blanked = true
	}

	return output, nil
}

func GetMuteStatus(request PJRequest) (se.MuteStatus, error) {

	response, err := HandleRequest(request)
	if err != nil {
		return se.MuteStatus{}, err
	}

	var output se.MuteStatus
	if strings.EqualFold(response.Response[0], "video and audio mute on") ||
		strings.EqualFold(response.Response[0], "audio mute on, video mute off") {
		output.Muted = true
	}

	return output, nil
}

func GetCurrentInput(request PJRequest) (se.Input, error) {

	response, err := HandleRequest(request)
	if err != nil {
		return se.Input{}, err
	}

	return se.Input{Input: response.Response[0]}, nil
}

func GetInputList(request PJRequest) (se.VideoList, error) {

	response, err := HandleRequest(request)
	if err != nil {
		return se.VideoList{}, err
	}

	var output se.VideoList
	for _, entry := range response.Response {

		output.Inputs = append(output.Inputs, se.Input{Input: entry})

	}

	return output, nil
}
