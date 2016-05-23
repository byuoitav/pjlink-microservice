package helpers

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"strings"
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

func HandleRequest(address, port, class, pwd, command, param string) (PjResponse, error) {
	//example values:
	//Address = "10.1.1.3"
	//Port = "4352"
	//Pwd = "magic123"
	//Class = "1"
	//Command = "POWR"
	//Param = "?"

	request := PjRequest{address, port, class, pwd, command, param}

	error1 := validateRequest(request)

	if error1 != nil {
		return PjResponse{}, error1
	} else {
		return parseResponse(sendRequest(request)), nil
	}
}

func validateRequest(request PjRequest) error {
	if len(request.Command) != 4 {
		return errors.New("your command doesn't have char length of 4")
	} else {
		return nil
	}
}

func parseResponse(response string) PjResponse {
	//if password is wrong, response will be 'PJLINK ERRA'
	fmt.Println(response)
	if strings.Contains(response, "ERRA") {
		return PjResponse{"0", "ERRA", "0"}
	} else {
		//example response: "%1POWR=0"
		fmt.Println(response)
		//params are class, command, and response code, respectively
		return PjResponse{response[1:2], response[2:6], response[7:len(response)]}
	}

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
