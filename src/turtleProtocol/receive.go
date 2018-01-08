package turtleProtocol

import (
	"bufio"
	"io"
	"log"
	"reflect"
	"turtleProtocol/msg"
)

func (s *Session)ReceiveLoop(){
	defer s.wgReceiver.Done()

	reader := bufio.NewReader(s.socket)

	packetSizeBytes := make([]byte, 4)

	for {
		_, err := io.ReadFull(reader, packetSizeBytes)
		if err != nil{log.Print(reflect.TypeOf(s).String() + err.Error());break}

		packetSize := fourBytesToInt(packetSizeBytes)
		packet := make([]byte, packetSize)

		_, err = io.ReadFull(reader, packet)
		if err != nil{log.Print(reflect.TypeOf(s).String() + err.Error());break}

		bytesLeft := packetSize

		numOfNonConfirmationMessages := 0

		for bytesLeft > 0{
			msgSizeBytes := packet[0:2]
			packet = packet[2:]

			msgSize := twoBytesToInt(msgSizeBytes)

			bytesLeft -= msgSize + len(msgSizeBytes)

			messageBytes := packet[0:msgSize]
			packet = packet[msgSize:]

			message, err := msg.FromBytes(messageBytes)

			if err != nil{
				log.Print("Receiver \"", s.name, "\" err ", err)
			} else if message.GetMessageType() == msg.OK {
				s.UnlockSending()
			} else {
				numOfNonConfirmationMessages++
				s.recv.OnReceive(message)
			}
		}

		if numOfNonConfirmationMessages > 0 {
			s.SendConfirmation()
		}
	}
}

func fourBytesToInt(size []byte)int{
	num := 0

	num += (int)(size[0])
	num += (int)(size[1]) * 256
	num += (int)(size[2]) * 256 * 256
	num += (int)(size[3]) * 256 * 256 * 256

	return num
}

func twoBytesToInt(size []byte)int{
	num := 0

	num += (int)(size[0])
	num += (int)(size[1]) * 256

	return num
}
