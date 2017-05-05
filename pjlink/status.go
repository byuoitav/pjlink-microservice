package pjlink

import "strings"

type PowerStatus struct {
	Power string `json:"power",omitempty`
}

type BlankedStatus struct {
	Blanked bool `json:"blanked",omitempty`
}

type MuteStatus struct {
	Mute bool `json:"muted",omitempty`
}

type VideoInput struct {
	Input string `json:"input",omitempty`
}

type InputList struct {
	Inputs []string `json:"inputs",omitempty`
}

func GetPowerStatus(request PJRequest) (PowerStatus, error) {

	response, err := HandleRequest(request)
	if err != nil {
		return PowerStatus{}, err
	}

	var output PowerStatus
	if strings.EqualFold(response.Response[0], "power-on (lamp on)") {
		output.Power = "on"
	} else {
		output.Power = "standby"
	}
	return output, nil
}

func GetBlankedStatus(request PJRequest) (BlankedStatus, error) {

	response, err := HandleRequest(request)
	if err != nil {
		return BlankedStatus{}, err
	}

	var output BlankedStatus
	if strings.EqualFold(response.Response[0], "video and audio mute on") ||
		strings.EqualFold(response.Response[0], "video mute on, audio mute off") {
		output.Blanked = true
	}

	return output, nil
}

func GetMuteStatus(request PJRequest) (MuteStatus, error) {

	response, err := HandleRequest(request)
	if err != nil {
		return MuteStatus{}, err
	}

	var output MuteStatus
	if strings.EqualFold(response.Response[0], "video and audio mute on") ||
		strings.EqualFold(response.Response[0], "audio mute on, video mute off") {
		output.Mute = true
	}

	return output, nil
}

func GetCurrentInput(request PJRequest) (VideoInput, error) {

	response, err := HandleRequest(request)
	if err != nil {
		return VideoInput{}, err
	}

	return VideoInput{Input: response.Response[0]}, nil
}

func GetInputList(request PJRequest) (InputList, error) {

	response, err := HandleRequest(request)
	if err != nil {
		return InputList{}, err
	}

	return InputList{Inputs: response.Response}, nil
}
