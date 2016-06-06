package helpers

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net"
	"strings"
)

type PJRequest struct {
	Address  string `json:"address"`
	Port     string `json:"port"`
	Class    string `json:"class"`
	Password string `json:"pwd"`
	Command  string `json:"command"`
	Param    string `json:"param"`
}

type PJResponse struct {
	Class   string `json:"class"`
	Command string `json:"command"`
	Code    string `json:"code"`
}

func PJLinkRequest(request PJRequest) (PJResponse, error) {
	return parseResponse(sendRequest(request)), nil
}

func parseResponse(response string) PJResponse {
	// If password is wrong, response will be 'PJLINK ERRA'
	if strings.Contains(response, "ERRA") {
		return PJResponse{"0", "ERRA", "0"}
	}

	return PJResponse{Class: response[1:2], Command: response[2:6], Code: response[7:len(response)]}
}

func sendRequest(request PJRequest) string {
	//pjlink always uses tcp
	protocol := "tcp"

	//establish TCP connection with PJLink device
	pjConn := connectToPJLink(protocol, request.Address, request.Port)

	//setup scanner
	scanner := bufio.NewScanner(pjConn)
	scanner.Split(bufio.ScanWords)

	sResponse := make([]string, 3)

	//grab initial response
	for i := 0; i < 3; i++ {
		scanner.Scan()
		sResponse[i] = scanner.Text()
	}

	//verify PJLink and correct class
	if !verifyPJLink(sResponse) {
		fmt.Println("Not a PJLINK class 1 connection!")
		//TODO handle not PJLink class 1
	}

	strCmd := generateCmd(sResponse[2], request.Pwd, request.Class,
		request.Command, request.Param)

	//test
	fmt.Println("sending: " + strCmd + "\n")

	//send command
	pjConn.Write([]byte(strCmd + "\r"))

	scanner.Scan()

	//if authentication failed, we received 'PJLINK ERRA', so return 'ERRA'
	if scanner.Text() == "PJLINK" {
		scanner.Scan()
	}
	pjConn.Close()

	return scanner.Text()
}

func generateCmd(seed, pjlinkPwd, pjlinkClass, pjlinkCmd, pjlinkParam string) string {
	return createEncryptedMsg(seed, pjlinkPwd) + "%" + pjlinkClass + pjlinkCmd +
		" " + pjlinkParam
}

func createEncryptedMsg(seed, pjlinkPwd string) string {

	//generate MD5
	data := []byte(seed + pjlinkPwd)
	hash := md5.Sum(data)

	//cast to string
	strHash := hex.EncodeToString(hash[:])

	return strHash
}

func verifyPJLink(sResponse []string) bool {

	if sResponse[0] != "PJLINK" {
		return false
	}

	if sResponse[1] != "1" {
		return false
	}

	return true
}

func connectToPJLink(protocolType, ip, port string) net.Conn {
	pjlinkConn, err := net.Dial(protocolType, ip+":"+port)
	if err != nil {
		// TODO handle
	}
	return pjlinkConn
}
