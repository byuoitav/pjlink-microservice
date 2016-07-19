package pjlink

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"
)

//HandleRawRequest is a wrapper function for all handling activity
//success: returns a populated PjResponse struct, nil error
//failure: returns empty PjResponse struct, error
func HandleRawRequest(request PJRequest) (PJResponse, error) {
	validateError := validateRawRequest(request)

	if validateError != nil { //malformed command, don't send
		return PJResponse{}, validateError
	} else { //send request and parse response into struct
		response, requestError := sendRawRequest(request)
		if requestError != nil {
			return PJResponse{}, requestError
		} else {
			return parseRawResponse(response)
		}
	}
}

//this function validates cmd length, before we send request.
//as of now this function only tests for 4 chars, which is pjlink standard cmd length
func validateRawRequest(request PJRequest) error {
	if len(request.Command) != 4 { // 4 characters is standard command length for PJLink
		return errors.New("Your command doesn't have character length of 4")
	}

	return nil
}

func sendRawRequest(request PJRequest) (string, error) {
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
	if !verifyPJLink(challenge) {
		// TODO: Handle not PJLink class 1
		return "", errors.New("Not a PJLINK class 1 connection")
	}

	stringCommand := generateCommand(challenge[2], request)

	//send command
	connection.Write([]byte(stringCommand + "\r"))
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

	connection, connectionError := net.DialTimeout(protocol, ip+":"+port,
		time.Duration(timeout)*time.Second)
	if connectionError != nil {
		return connection, errors.New("failed to establish a connection with " +
			"pjlink device. error msg: " + connectionError.Error())
	}
	return connection, connectionError
}

//handle and parse response
//returns a populated PJResponse struct
func parseRawResponse(response string) (PJResponse, error) {
	// If password is wrong, response will be 'PJLINK ERRA'
	if strings.Contains(response, "ERRA") { //if authentication succeeded
		return PJResponse{}, errors.New("Incorrect password")
		//example response: "%1POWR=0"
		//returned params are class, command, and response code(s), respectively
	}

	tokens := strings.Split(response, " ")
	fmt.Printf("tokens: %v\n", tokens)

	token0 := tokens[0]
	param1 := []string{token0[7:len(token0)]}
	paramsN := tokens[1:len(tokens)]
	params := append(param1, paramsN...)

	return PJResponse{Class: token0[1:2], Command: token0[2:6], Response: params}, nil
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
