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

		//msgOK
		if packetSize == 4{
			msgSizeBytes := packet[0:2]
			packet = packet[2:]

			msgSize := twoBytesToInt(msgSizeBytes)
			if msgSize == 2{
				content := packet[0:msgSize]
				message, err := msg.FromBytes(content)
				if err == nil && message.GetMessageType() == msg.OK{
					s.UnlockSending()
				}
			}
		} else {
			bytesLeft := packetSize

			for bytesLeft > 0{
				msgSizeBytes := packet[0:2]
				packet = packet[2:]

				msgSize := twoBytesToInt(msgSizeBytes)

				bytesLeft -= msgSize + len(msgSizeBytes)

				message := packet[0:msgSize]
				packet = packet[msgSize:]

				s.recv.OnReceive(message)
			}
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
