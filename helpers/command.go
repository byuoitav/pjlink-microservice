package helpers

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net"
)

type PjRequest struct {
	Address string `json:address`
	Port    string `json:port`
	Class   string `json:class`
	Pwd     string `json:pwd`
	Command string `json:command`
	Param   string `json:param`
}

type PjResponse struct {
	Class   string `json:"class"`
	Command string `json:"command"`
	Code    string `json:"code"`
}

func Test() (string, error) {
	return "check yo head!", nil
}

func PjlinkRequest(address, port, class, pwd, command, param string) (PjResponse, error) {
	//example values:
	//address = "10.1.1.3"
	//port = "4352"
	//pwd = "magic123"
	//class = "1"
	//command = "POWR"
	//param = "?"

	request := PjRequest{address, port, class, pwd, command, param}

	return parseResponse(sendRequest(request)), nil
}

func parseResponse(response string) PjResponse {
	//example response: "%1POWR=0"
	fmt.Println(response)
	//params are class, command, and response code, respectively
	parsedResponse := PjResponse{response[1:2], response[2:6], response[7:len(response)]}
	return parsedResponse
}

func sendRequest(request PjRequest) string {
	//pjlink always uses tcp
	protocol := "tcp"

	//establish TCP connection with PJLink device
	pjConn := connectToPjlink(protocol, request.Address, request.Port)

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
	if !verifyPjlink(sResponse) {
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

func verifyPjlink(sResponse []string) bool {

	if sResponse[0] != "PJLINK" {
		return false
	}

	if sResponse[1] != "1" {
		return false
	}

	return true
}

func connectToPjlink(protocolType, ip, port string) net.Conn {
	pjlinkConn, err := net.Dial(protocolType, ip+":"+port)
	if err != nil {
		// TODO handle
	}
	return pjlinkConn
}
