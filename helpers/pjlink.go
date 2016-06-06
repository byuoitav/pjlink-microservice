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

//wrapper function for all handling activity
//success: returns a populated PjResponse struct, nil error
//failure: returns empty PjResponse struct, error
func handleRequest(request PJRequest) (PJResponse, error) {

	validateError = validateRequest(request)

	if validateError != nil { //malformed command, don't send
		return PJResponse{}, validateError
	} else { //send request and parse response into struct
		response, requestError := sendRequest(request)
		if requestError != nil {
			return PJResponse{}, requestError
		} else {
			return parseResponse(response), nil
		}
	}
}

//this function validates cmd length, before we send request.
//as of now this function only tests for 4 chars, which is pjlink standard cmd length
func validateRequest(request PJRequest) error {
	if len(request.Command) != 4 { // 4 characters is standard command length for PJLink
		return errors.New("Your command doesn't have character length of 4")
	} else { //authentication succeedded

		return nil
	}
}

func sendRequest(request PJRequest) (string, error) {
	//establish TCP connection with PJLink device
	connection, connectionError := connectToPJLink(request.Address, request.Port)

	if connectionError != nil {
		return "", connectionError
	}

	// Define a split function that separates on carriage return (i.e '\r').
	onCarriageReturn := func(data []byte, atEOF bool) (advance int, token []byte,
		err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == '\r' {
				return i + 1, data[:i], nil
			}
		}
		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens
		// after this but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	}

	//setup scanner
	scanner := bufio.NewScanner(connection)
	scanner.Split(onCarriageReturn)
	scanner.Scan() //grab a line
	challenge := strings.Split(scanner.Text(), " ")

	//verify PJLink and correct class
	if !verifyPJLink(response) {
		// TODO: Handle not PJLink class 1
		return "", errors.New("Not a PJLINK class 1 connection")
	}

	stringCommand := generateCommand(challenge[2], request)

	//send command
	connection.Write([]byte(command + "\r"))
	scanner.Scan() //grab response line

	connection.Close()

	return scanner.Text(), nil
}

//attempts to establish a TCP socket with the specified IP:port
//success: returns populated pjlinkConn struct and nil error
//failure: returns empty pjlinkConn and error
func connectToPJLink(ip, port string) (net.Conn, error) {
	protocol := "tcp" //PJLink always uses TCP
	timeout := 5      //represents seconds

	connection, connectionError := net.DialTimeout(protocolType, ip+":"+port,
		time.Duration(timeout)*time.Second)
	if connectionError != nil {
		return connection, errors.New("failed to establish a connection with " +
			"pjlink device. error msg: " + connectionError.Error())
	}
	return pjlinkConn, connectionError
}

//handle and parse response
//returns a populated PJResponse struct
func parseResponse(response string) (PJResponse, error) {
	// If password is wrong, response will be 'PJLINK ERRA'
	if strings.Contains(response, "ERRA") {
		return PJResponse{}, errors.New("Incorrect password")
	} else { //authentication succeeded
		//example response: "%1POWR=0"
		//returned params are class, command, and response code(s), respectively

		tokens := strings.Split(response, " ")
		fmt.Printf("tokens: %v", tokens)

		token0 := tokens[0]
		class := token0[1:2]
		command := token0[2:6]
		param1 := token0[7:len(token0)]
		params := []string{param1, tokens[1:len(tokens)]}

		return PJResponse{Class: response[1:2], Command: response[2:6], Response: response[7:len(response)]}, nil
	}

}

//returns PJLink command string
func generateCommand(seed string, request PJRequest) string {
	return createEncryptedMessage(seed, request.Password) + "%" +
		request.Class + request.Command + " " + request.Parameter
}

//generates a hash given seed and password
//returns string hash
func createEncryptedMessage(seed, password string) string {
	//generate MD5
	data := []byte(seed + password)
	hash := md5.Sum(data)

	//cast to string
	stringHash := hex.EncodeToString(hash[:])

	return stringHash
}

//verify we receive a pjlink class 1 challenge
//success: returns true
//failure: returns false
func verifyPJLink(response []string) bool {
	if response[0] != "PJLINK" {
		return false
	}

	if response[1] != "1" {
		return false
	}

	return true
}
