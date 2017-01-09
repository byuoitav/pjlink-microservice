package pjlink

import (
	"errors"
	"log"
	"strconv"
)

//HandleRequest takes a PJRequest in human readable format
//success: returns a human readable PJResponse
//fail: returns an empty PJResponse, and accompanying error message
func HandleRequest(request PJRequest) (PJResponse, error) {
	//first validate request before sending
	validateRequestError := validateHumanRequest(request)
	if validateRequestError != nil {
		return PJResponse{}, errors.New("Not a valid request: " + validateRequestError.Error())
	}

	//convert raw PJLink request
	rawResponse, handleRawRequestError := HandleRawRequest(
		convertHumanRequestToRawRequest(request))

	if handleRawRequestError != nil {
		return PJResponse{}, handleRawRequestError
	} else {
		response, error := convertRawResponseToHumanResponse(
			rawResponse, request.Parameter)
		if error != nil {
			return PJResponse{}, error
		} else {
			return response, nil
		}
	}
}

func convertHumanRequestToRawRequest(request PJRequest) PJRequest {
	rawPJRequest := PJRequest{
		Address:  request.Address,
		Port:     request.Port,
		Password: request.Password,
		Class:    request.Class,
		Command:  HumanToRawCommands[request.Command],
	}

	switch request.Command {
	case "power":
		rawPJRequest.Parameter = PowerRequests[request.Parameter]
	case "input-list":
		rawPJRequest.Parameter = InputListRequests[request.Parameter]
	case "input":
		rawPJRequest.Parameter = InputRequests[request.Parameter]
	case "av-mute":
		rawPJRequest.Parameter = AVMuteRequests[request.Parameter]
	case "error-status":
		rawPJRequest.Parameter = ErrorStatusRequests[request.Parameter]
	case "lamp":
		rawPJRequest.Parameter = LampRequests[request.Parameter]
	case "name":
		rawPJRequest.Parameter = NameRequests[request.Parameter]
	case "manufacturer":
		rawPJRequest.Parameter = ManufacturerRequests[request.Parameter]
	case "model":
		rawPJRequest.Parameter = ModelRequests[request.Parameter]
	case "version":
		rawPJRequest.Parameter = VersionRequests[request.Parameter]
	}

	log.Printf("humanRequest: %+v", request)
	log.Printf("rawRequest:   %+v", rawPJRequest)

	return rawPJRequest
}

func convertRawResponseToHumanResponse(rawResponse PJResponse,
	requestParameter string) (PJResponse, error) {
	var convertError error

	log.Printf("rawResponse: %+v", rawResponse)

	response := PJResponse{
		Class:    rawResponse.Class,
		Command:  RawToHumanCommands[rawResponse.Command],
		Response: make([]string, len(rawResponse.Response)),
	}

	switch response.Command {
	case "power":
		if requestParameter == "query" {
			response.Response[0] = PowerQueryResponses[rawResponse.Response[0]]
			if response.Response[0] == "" {
				convertError = errors.New("unknown power response")
			}
		} else {
			response.Response[0] = PowerResponses[rawResponse.Response[0]]
		}
	case "input-list":
		if (rawResponse.Response[0] == "ERR3") || (rawResponse.Response[0] == "ERR4") {
			response.Response[0] = InputListQueryResponses[rawResponse.Response[0]]
		}
		response.Response = interpretInputListInputs(rawResponse.Response)
	case "input":
		if requestParameter == "query" {
			response.Response[0] = InputQueryResponses[rawResponse.Response[0]]
		} else {
			response.Response[0] = InputResponses[rawResponse.Response[0]]
		}
	case "av-mute":
		if requestParameter == "query" {
			response.Response[0] = AVMuteQueryResponses[rawResponse.Response[0]]
		} else { //command
			response.Response[0] = AVMuteResponses[rawResponse.Response[0]]
		}
	case "error-status":
		if (rawResponse.Response[0] == "ERR3") || (rawResponse.Response[0] == "ERR4") {
			response.Response[0] = AVMuteQueryResponses[rawResponse.Response[0]]
		} else {
			humanErrorResponses := interpretErrorStatusResponse(rawResponse.Response)
			log.Println(humanErrorResponses)
			response.Response = humanErrorResponses
		}

	case "lamp":
		if (rawResponse.Response[0] == "ERR3") || (rawResponse.Response[0] == "ERR4") {
			response.Response[0] = LampQueryResponses[rawResponse.Response[0]]
		} else {
			response.Response = interpretLampQueryResponse(rawResponse.Response)
		}
	case "name":
		if (rawResponse.Response[0] == "ERR3") || (rawResponse.Response[0] == "ERR4") {
			response.Response[0] = NameQueryResponses[rawResponse.Response[0]]
		} else {
			response.Response[0] = rawResponse.Response[0]
		}
	case "manufacturer":
		if (rawResponse.Response[0] == "ERR3") || (rawResponse.Response[0] == "ERR4") {
			response.Response[0] = ManufacturerQueryResponses[rawResponse.Response[0]]
		} else {
			response.Response[0] = rawResponse.Response[0]
		}
	case "model":
		if (rawResponse.Response[0] == "ERR3") || (rawResponse.Response[0] == "ERR4") {
			response.Response[0] = ModelQueryResponses[rawResponse.Response[0]]
		} else {
			response.Response[0] = rawResponse.Response[0]
		}
	case "version":
		if (rawResponse.Response[0] == "ERR3") || (rawResponse.Response[0] == "ERR4") {
			response.Response[0] = VersionQueryResponses[rawResponse.Response[0]]
		} else {
			response.Response = rawResponse.Response
		}
	}

	log.Printf("humanResponse: %+v", response)

	return response, convertError
}

func odd(number int) bool {
	return number%2 != 0
}

func interpretLampQueryResponse(rawLampResponse []string) []string {
	sets := len(rawLampResponse) / 2
	humanLampResponses := make([]string, sets)
	for index := range humanLampResponses {
		humanLampResponses[index] = "Lamp: " + strconv.Itoa(index+1) +
			", hours: " + rawLampResponse[index*2] + ", state: " +
			" " + LampStateResponses[rawLampResponse[(index*2)+1]]
	}

	return humanLampResponses
}

func interpretErrorStatusResponse(rawResponse []string) []string {
	log.Println(rawResponse)

	raw := rawResponse[0]
	rawValues := make([]string, 6)

	for index := range rawValues {
		rawValues[index] = raw[index : index+1]
	}

	log.Println(rawValues)

	humanCodes := make([]string, 6)
	for index := range humanCodes {
		switch rawValues[index] {
		case "0":
			humanCodes[index] = "OK"
		case "1":
			humanCodes[index] = "Warning"
		case "2":
			humanCodes[index] = "Error"
		}
	}

	humanResponses := []string{
		"fan: " + humanCodes[0],
		"lamp: " + humanCodes[1],
		"temperature: " + humanCodes[2],
		"cover open: " + humanCodes[3],
		"filter: " + humanCodes[4],
		"other: " + humanCodes[5],
	}

	return humanResponses
}

func interpretInputListInputs(rawResponses []string) []string {
	humanResponses := make([]string, len(rawResponses))

	for index, element := range rawResponses {
		humanResponses[index] = RawToHumanInputs[element]
	}

	return humanResponses
}

func validateHumanRequest(request PJRequest) error {
	//class 1?
	if request.Class != "1" {
		return errors.New("only PJLink class 1 supported")
	}

	//valid PJLink class 1 command?
	if HumanToRawCommands[request.Command] == "" {
		return errors.New("'" + request.Command + "' is not a valid PJLink command.")
	}

	//valid parameter for given command?
	validateCommandParameterError := validateCommandParameter(request)
	if validateCommandParameterError != nil {
		return validateCommandParameterError
	}

	//all is well, so return nil error
	return nil
}

func validateCommandParameter(request PJRequest) error {
	var badParameter bool = false

	switch request.Command {
	case "power":
		if PowerRequests[request.Parameter] == "" {
			badParameter = true
		}
	case "input-list":
		if InputListRequests[request.Parameter] == "" {
			badParameter = true
		}
	case "input":
		if InputRequests[request.Parameter] == "" {
			badParameter = true
		}
	case "av-mute":
		if AVMuteRequests[request.Parameter] == "" {
			badParameter = true
		}
	case "error-status":
		if ErrorStatusRequests[request.Parameter] == "" {
			badParameter = true
		}
	case "lamp":
		if ErrorStatusRequests[request.Parameter] == "" {
			badParameter = true
		}
	case "name":
		if ErrorStatusRequests[request.Parameter] == "" {
			badParameter = true
		}
	case "manufacturer":
		if ErrorStatusRequests[request.Parameter] == "" {
			badParameter = true
		}
	case "model":
		if ErrorStatusRequests[request.Parameter] == "" {
			badParameter = true
		}
	case "version":
		if ErrorStatusRequests[request.Parameter] == "" {
			badParameter = true
		}
	}

	if badParameter {
		return errors.New("'" + request.Parameter +
			"' is not a valid parameter for PJLink command '" +
			request.Command + "'.")
	}

	return nil
}
