package main

import (
	"fmt"
	"net"
	"log"
	"encoding/binary"
	"regexp"
	"connectionListener"
)




func addSizeToMessage(message []byte) (outputMessage []byte) {
	messageSizeArray := make([]byte, 4)
	messageSize := uint32(cap(message))
	binary.LittleEndian.PutUint32(messageSizeArray[0:], messageSize)
	outputMessage = append(messageSizeArray, message...)
	return outputMessage
}

func sendTo(writeTo string,  message []byte) {

	fmt.Println(writeTo)
	connection, err := net.Dial("tcp", writeTo )
	if err != nil {
		log.Fatal(err)
	}
	connection.Write(message)
}

func getIp(message string) (ipAddress string, newMessage []byte)  {
	re := regexp.MustCompile("(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]):[0-9]+ ")
	ipAddress = re.FindString(message)
	fmt.Println(ipAddress)
	reg := regexp.MustCompile(ipAddress)
	message = reg.ReplaceAllString(message, "")
	newMessage = []byte(message)

	return ipAddress, newMessage
}






func main() {
	fmt.Println("Hello")
	clnc, err:= connectionListener.NewConnectionListener("4000", nil)//TODO handler
	if err != nil {
		clnc.Loop()
	}

	clns, err := connectionListener.NewConnectionListener("4001", nil)//TODO handler
	if err != nil {
		clns.Loop()
	}
}