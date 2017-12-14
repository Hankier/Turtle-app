package textReceiver

import (
	"fmt"
	"strings"
)

const (
	MAX_TYPE_LEN = 10
)

type TextReceiverImpl struct{

}

func (*TextReceiverImpl)Print(from string, text string){
	paddinglen := 10 - len(from)
	if paddinglen < 0 {paddinglen = 0}
	fmt.Println("type: " + from + strings.Repeat(" ", paddinglen) + " - message: " + text)
}