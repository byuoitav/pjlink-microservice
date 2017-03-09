package pjlink

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertHumanRequestToRawRequest_power(test *testing.T) {
	var humanRequest PJRequest
	var rawRequest PJRequest

	humanRequest = helperCreateHumanRequest("power", "query")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)

	assert.Equal(test, rawRequest.Command, "POWR")
	assert.NotEqual(test, rawRequest.Command, "POWr")
	assert.Equal(test, rawRequest.Parameter, "?")
	assert.NotEqual(test, rawRequest.Parameter, "query")

	humanRequest = helperCreateHumanRequest("power", "power-on")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "POWR")
	assert.NotEqual(test, rawRequest.Command, "PoWr")
	assert.Equal(test, rawRequest.Parameter, "1")
	assert.NotEqual(test, rawRequest.Parameter, "2")

	humanRequest = helperCreateHumanRequest("power", "power-off")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "POWR")
	assert.Equal(test, rawRequest.Parameter, "0")
}

func TestConvertHumanRequestToRawRequest_inputlist(test *testing.T) {
	humanRequest := helperCreateHumanRequest("input-list", "query")
	rawRequest := convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INST")
	assert.Equal(test, rawRequest.Parameter, "?")
}

func TestConvertHumanRequestToRawRequest_input(test *testing.T) {
	var humanRequest PJRequest
	var rawRequest PJRequest

	humanRequest = helperCreateHumanRequest("input", "query")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "?")

	humanRequest = helperCreateHumanRequest("input", "rgb1")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "11")

	humanRequest = helperCreateHumanRequest("input", "rgb2")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "12")

	humanRequest = helperCreateHumanRequest("input", "rgb3")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "13")

	humanRequest = helperCreateHumanRequest("input", "rgb4")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "14")

	humanRequest = helperCreateHumanRequest("input", "rgb5")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "15")

	humanRequest = helperCreateHumanRequest("input", "rgb6")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "16")

	humanRequest = helperCreateHumanRequest("input", "rgb7")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "17")

	humanRequest = helperCreateHumanRequest("input", "rgb8")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "18")

	humanRequest = helperCreateHumanRequest("input", "rgb9")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "19")

	humanRequest = helperCreateHumanRequest("input", "video1")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "21")

	humanRequest = helperCreateHumanRequest("input", "video2")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "22")

	humanRequest = helperCreateHumanRequest("input", "video3")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "23")

	humanRequest = helperCreateHumanRequest("input", "video4")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "24")

	humanRequest = helperCreateHumanRequest("input", "video5")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "25")

	humanRequest = helperCreateHumanRequest("input", "video6")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "26")

	humanRequest = helperCreateHumanRequest("input", "video7")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "27")

	humanRequest = helperCreateHumanRequest("input", "video8")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "28")

	humanRequest = helperCreateHumanRequest("input", "video9")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "29")

	humanRequest = helperCreateHumanRequest("input", "digital1")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "31")

	humanRequest = helperCreateHumanRequest("input", "digital2")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "32")

	humanRequest = helperCreateHumanRequest("input", "digital3")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "33")

	humanRequest = helperCreateHumanRequest("input", "digital4")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "34")

	humanRequest = helperCreateHumanRequest("input", "digital5")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "35")

	humanRequest = helperCreateHumanRequest("input", "digital6")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "36")

	humanRequest = helperCreateHumanRequest("input", "digital7")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "37")

	humanRequest = helperCreateHumanRequest("input", "digital8")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "38")

	humanRequest = helperCreateHumanRequest("input", "digital9")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "39")

	humanRequest = helperCreateHumanRequest("input", "storage1")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "41")

	humanRequest = helperCreateHumanRequest("input", "storage2")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "42")

	humanRequest = helperCreateHumanRequest("input", "storage3")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "43")

	humanRequest = helperCreateHumanRequest("input", "storage4")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "44")

	humanRequest = helperCreateHumanRequest("input", "storage5")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "45")

	humanRequest = helperCreateHumanRequest("input", "storage6")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "46")

	humanRequest = helperCreateHumanRequest("input", "storage7")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "47")

	humanRequest = helperCreateHumanRequest("input", "storage8")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "48")

	humanRequest = helperCreateHumanRequest("input", "storage9")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "49")

	humanRequest = helperCreateHumanRequest("input", "network1")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "51")

	humanRequest = helperCreateHumanRequest("input", "network2")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "52")

	humanRequest = helperCreateHumanRequest("input", "network3")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "53")

	humanRequest = helperCreateHumanRequest("input", "network4")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "54")

	humanRequest = helperCreateHumanRequest("input", "network5")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "55")

	humanRequest = helperCreateHumanRequest("input", "network6")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "56")

	humanRequest = helperCreateHumanRequest("input", "network7")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "57")

	humanRequest = helperCreateHumanRequest("input", "network8")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "58")

	humanRequest = helperCreateHumanRequest("input", "network9")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INPT")
	assert.Equal(test, rawRequest.Parameter, "59")
}

func TestConvertHumanRequestToRawRequest_avmute(test *testing.T) {
	var humanRequest PJRequest
	var rawRequest PJRequest

	humanRequest = helperCreateHumanRequest("av-mute", "query")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "AVMT")
	assert.Equal(test, rawRequest.Parameter, "?")

	humanRequest = helperCreateHumanRequest("av-mute", "video-mute-on")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "AVMT")
	assert.Equal(test, rawRequest.Parameter, "11")

	humanRequest = helperCreateHumanRequest("av-mute", "video-mute-off")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "AVMT")
	assert.Equal(test, rawRequest.Parameter, "10")

	humanRequest = helperCreateHumanRequest("av-mute", "audio-mute-on")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "AVMT")
	assert.Equal(test, rawRequest.Parameter, "21")

	humanRequest = helperCreateHumanRequest("av-mute", "audio-mute-off")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "AVMT")
	assert.Equal(test, rawRequest.Parameter, "20")

	humanRequest = helperCreateHumanRequest("av-mute", "av-mute-on")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "AVMT")
	assert.Equal(test, rawRequest.Parameter, "31")

	humanRequest = helperCreateHumanRequest("av-mute", "av-mute-off")
	rawRequest = convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "AVMT")
	assert.Equal(test, rawRequest.Parameter, "30")
}

func TestConvertHumanRequestToRawRequest_errorstatus(test *testing.T) {
	humanRequest := helperCreateHumanRequest("error-status", "query")
	rawRequest := convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "ERST")
	assert.Equal(test, rawRequest.Parameter, "?")
}

func TestConvertHumanRequestToRawRequest_lamp(test *testing.T) {
	humanRequest := helperCreateHumanRequest("lamp", "query")
	rawRequest := convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "LAMP")
	assert.Equal(test, rawRequest.Parameter, "?")
}

func TestConvertHumanRequestToRawRequest_name(test *testing.T) {
	humanRequest := helperCreateHumanRequest("name", "query")
	rawRequest := convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "NAME")
	assert.Equal(test, rawRequest.Parameter, "?")
}

func TestConvertHumanRequestToRawRequest_manufacturer(test *testing.T) {
	humanRequest := helperCreateHumanRequest("manufacturer", "query")
	rawRequest := convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INF1")
	assert.Equal(test, rawRequest.Parameter, "?")
}

func TestConvertHumanRequestToRawRequest_model(test *testing.T) {
	humanRequest := helperCreateHumanRequest("model", "query")
	rawRequest := convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INF2")
	assert.Equal(test, rawRequest.Parameter, "?")
}

func TestConvertHumanRequestToRawRequest_version(test *testing.T) {
	humanRequest := helperCreateHumanRequest("version", "query")
	rawRequest := convertHumanRequestToRawRequest(humanRequest)
	assert.Equal(test, rawRequest.Command, "INFO")
	assert.Equal(test, rawRequest.Parameter, "?")
}

func TestConvertRawResponseToHumanResponse_powerQueryResponses(test *testing.T) {
	var rawResponse PJResponse
	var humanResponse PJResponse
	var convertError error

	rawResponse = createRawResponse("POWR", []string{"0"})
	humanResponse, convertError = convertRawResponseToHumanResponse(rawResponse, "query")
	if convertError != nil {
		assert.Fail(test, "error isn't nil")
	}
	assert.Equal(test, humanResponse.Command, "power")
	assert.Equal(test, humanResponse.Response[0], "power-off (standby)")

	rawResponse = createRawResponse("POWR", []string{"1"})
	humanResponse, convertError = convertRawResponseToHumanResponse(rawResponse, "query")
	if convertError != nil {
		assert.Fail(test, "error isn't nil")
	}
	assert.Equal(test, humanResponse.Command, "power")
	assert.Equal(test, humanResponse.Response[0], "power-on (lamp on)")

	rawResponse = createRawResponse("POWR", []string{"2"})
	humanResponse, convertError = convertRawResponseToHumanResponse(rawResponse, "query")
	if convertError != nil {
		assert.Fail(test, "error isn't nil")
	}
	assert.Equal(test, humanResponse.Command, "power")
	assert.Equal(test, humanResponse.Response[0], "cooling")

	rawResponse = createRawResponse("POWR", []string{"3"})
	humanResponse, convertError = convertRawResponseToHumanResponse(rawResponse, "query")
	if convertError != nil {
		assert.Fail(test, "error isn't nil")
	}
	assert.Equal(test, humanResponse.Command, "power")
	assert.Equal(test, humanResponse.Response[0], "warm-up")

	rawResponse = createRawResponse("POWR", []string{"ERR3"})
	humanResponse, convertError = convertRawResponseToHumanResponse(rawResponse, "query")
	if convertError != nil {
		assert.Fail(test, "error isn't nil")
	}
	assert.Equal(test, humanResponse.Command, "power")
	assert.Equal(test, humanResponse.Response[0], "unavailable time")

	rawResponse = createRawResponse("POWR", []string{"ERR4"})
	humanResponse, convertError = convertRawResponseToHumanResponse(rawResponse, "query")
	if convertError != nil {
		assert.Fail(test, "error isn't nil")
	}
	assert.Equal(test, humanResponse.Command, "power")
	assert.Equal(test, humanResponse.Response[0], "device failure")

	rawResponse = createRawResponse("POWR", []string{"ERR5"})
	humanResponse, convertError = convertRawResponseToHumanResponse(rawResponse, "query")
	if convertError == nil {
		assert.Fail(test, "failed")
	}
}

//helper functions
func helperCreateHumanRequest(humanCommand string, humanParameter string) PJRequest {
	humanPJRequest := PJRequest{
		Address:   "n/a",
		Port:      "n/a",
		Password:  "n/a",
		Command:   humanCommand,
		Parameter: humanParameter,
	}
	return humanPJRequest
}

func createRawResponse(rawCommand string, rawParameter []string) PJResponse {
	rawPJResponse := PJResponse{
		Class:    "n/a",
		Command:  rawCommand,
		Response: rawParameter,
	}
	return rawPJResponse
}
