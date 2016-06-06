package helpers

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"net"
	"strings"
	"time"
)

type PJRequest struct {
	Address   string `json:"address"`
	Port      string `json:"port"`
	Class     string `json:"class"`
	Password  string `json:"password"`
	Command   string `json:"command"`
	Parameter string `json:"parameter"`
}

type PJResponse struct {
	Class    string `json:"class"`
	Command  string `json:"command"`
	Response string `json:"response"`
}

func PJLinkRequest(request PJRequest) (PJResponse, error) {
	response, err := sendRequest(request)
	if err != nil {
		return PJResponse{}, err
	}

	err = validateRequest(request)
	if err != nil {
		return PJResponse{}, err
	}

	parsedResponse, err := parseResponse(response)
	if err != nil {
		return PJResponse{}, err
	}

	return parsedResponse, nil
}

func validateRequest(request PJRequest) error {
	if len(request.Command) != 4 { // 4 characters is standard command length for PJLink
		return errors.New("Your command doesn't have character length of 4")
	}

	return nil
}

func parseResponse(response string) (PJResponse, error) {
	// If password is wrong, response will be 'PJLINK ERRA'
	if strings.Contains(response, "ERRA") {
		return PJResponse{}, errors.New("Incorrect password")
	}

	return PJResponse{Class: response[1:2], Command: response[2:6], Response: response[7:len(response)]}, nil
}

func sendRequest(request PJRequest) (string, error) {
	//establish TCP connection with PJLink device
	timeout := 5 // In seconds
	connection, err := net.DialTimeout("tcp", request.Address+":"+request.Port, time.Duration(timeout)*time.Second)
	if err != nil {
		return "", err
	}

	//setup scanner
	scanner := bufio.NewScanner(connection)
	scanner.Split(bufio.ScanWords)

	response := make([]string, 3)

	//grab initial response
	for i := 0; i < 3; i++ {
		scanner.Scan()
		response[i] = scanner.Text()
	}

	//verify PJLink and correct class
	if !verifyPJLink(response) {
		// TODO: Handle not PJLink class 1
		return "", errors.New("Not a PJLINK class 1 connection")
	}

	command := generateCommand(response[2], request)

	//send command
	connection.Write([]byte(command + "\r"))
	scanner.Scan()

	connection.Close()

	return scanner.Text(), nil
}

func generateCommand(seed string, request PJRequest) string {
	return createEncryptedMessage(seed, request.Password) + "%" + request.Class + request.Command + " " + request.Parameter
}

func createEncryptedMessage(seed, password string) string {
	//generate MD5
	data := []byte(seed + password)
	hash := md5.Sum(data)

	//cast to string
	stringHash := hex.EncodeToString(hash[:])

	return stringHash
}

func verifyPJLink(response []string) bool {
	if response[0] != "PJLINK" {
		return false
	}

	if response[1] != "1" {
		return false
	}

	return true
}
