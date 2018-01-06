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

func (*TextReceiverImpl)Print(cmd string, text string){
	const maxcmd = 20
	if len(cmd) >= maxcmd{
		cmd = cmd[0:maxcmd]
	}
	paddinglen := maxcmd - len(cmd)
	if paddinglen < 0 {paddinglen = 0}
	fmt.Println(cmd + strings.Repeat(" ", paddinglen) + ": " + text)
}