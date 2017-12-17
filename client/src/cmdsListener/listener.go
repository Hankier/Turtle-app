package cmdsListener

import (
	"textReceiver"
	"bufio"
	"os"
	"strings"
	"strconv"
	"client/ui"
)

type Listener struct{
	ui       ui.UserInterface
	textrecv textReceiver.TextReceiver
	run      bool
	lastCmd  string
}

func New(ui ui.UserInterface, textrecv textReceiver.TextReceiver)(*Listener){
	cmdl := new(Listener)
	cmdl.ui = ui
	cmdl.textrecv = textrecv
	cmdl.run = true
	cmdl.lastCmd = ""

	return cmdl
}

func (cmdl *Listener)Listen(){
	reader := bufio.NewReader(os.Stdin)

	for cmdl.run{
		text, err := reader.ReadString('\n')
		text = text[:len(text) - 1] //cut newline
		if err != nil{
			cmdl.run = false
		} else {
			cmdl.execCmd(text)
		}
	}
}

func (cmdl *Listener)execCmd(cmd string){
	//arrow up
	if cmd == (string)([]byte{27, 91, 65}){
		cmd = cmdl.lastCmd
	}

	cmdl.lastCmd = cmd

	cmds := strings.Fields(cmd)
	cmdl.textrecv.Print("command", cmd)

	if len(cmds) > 0 {
		switch cmds[0] {
		case "get":
			if len(cmds) > 1 {
				switch cmds[1] {
				case "path":
					cmdl.textrecv.Print("path", strings.Join(cmdl.ui.GetCurrentPath(), " "))
				case "servers":
					cmdl.textrecv.Print("servers", strings.Join(cmdl.ui.GetServerList(), " "))
				}
			} else {
				cmdl.textrecv.Print("error", "usage: get path, get servers")
			}
			break
		case "connect":
			if len(cmds) > 1 {
				err := cmdl.ui.ConnectToServer(cmds[1])
				if err != nil {
					cmdl.textrecv.Print("Error: ", "Wrong server")
				} else {
					cmdl.textrecv.Print("Connecting to server ", cmds[1])
				}
			} else {
				cmdl.textrecv.Print("error", "usage: connect serverName")
			}
			break
		case "new":
			if len(cmds) > 1 {
				switch cmds[1] {
				case "convo":
					if len(cmds) > 3 {
						err := cmdl.ui.CreateConversation(cmds[2], cmds[3])
						if err != nil {
							cmdl.textrecv.Print("Error: ", err.Error())
						} else {
							cmdl.textrecv.Print("Created conversation ", cmds[2]+" "+cmds[3])
						}
					} else {
						cmdl.textrecv.Print("error", "usage: new convo clientName serverName")
					}
				case "path":
					if len(cmds) > 2 {
						length, err := strconv.Atoi(cmds[2])
						if err != nil {
							cmdl.textrecv.Print("Error: ", err.Error())
						} else {
							path, err := cmdl.ui.ChooseNewPath(length)
							if err != nil{
								cmdl.textrecv.Print("new path", err.Error())
							} else {
								cmdl.textrecv.Print("new path", strings.Join(path, " "))
							}
						}
					} else {
						cmdl.textrecv.Print("error", "usage: new path length")
					}
				}
			} else {
				cmdl.textrecv.Print("error", "usage: new convo clientName serverName, new path length")
			}
			break
		case "send":
			if len(cmds) > 2 {
				receiverServer := cmds[1]
				receiver := cmds[2]
				message := strings.Join(cmds[3:], " ")
				err := cmdl.ui.SendTo(receiverServer, receiver, message)
				if err != nil {
					cmdl.textrecv.Print("Error: ", err.Error())
				} else {
					cmdl.textrecv.Print("Message sent to ", cmds[1]+" "+cmds[2])
				}
			} else {
				cmdl.textrecv.Print("error", "usage: send receiverServer receiver message")
			}
			break
		case "exit":
			cmdl.run = false
			break
		default:
			cmdl.textrecv.Print("error", "avaiable cmds: get connect new send exit")
		}
	}
}