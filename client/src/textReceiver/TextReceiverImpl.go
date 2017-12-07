package textReceiver

import "log"

type TextReceiverImpl struct{

}

func (*TextReceiverImpl)Print(from string, text string){
	log.Print("Received from: " + from + " message: " + text)
}