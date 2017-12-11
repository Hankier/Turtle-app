package commandsListener

import (
	"userInterface"
	"textReceiver"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type CommandsListener struct{
	ui       userInterface.UserInterface
	textrecv textReceiver.TextReceiver
	run      bool
}

func New(ui userInterface.UserInterface, textrecv textReceiver.TextReceiver)(*CommandsListener){
	cmdl := new(CommandsListener)
	cmdl.ui = ui
	cmdl.textrecv = textrecv
	cmdl.run = true

	return cmdl
}

func (cmdl *CommandsListener)Listen(){
	reader := bufio.NewReader(os.Stdin)

	for cmdl.run{
		text, err := reader.ReadString('\n')
		if err != nil{
			cmdl.run = false
		} else {
			cmdl.execCmd(text)
		}
	}
}

func (cmdl *CommandsListener)execCmd(cmd string){
	cmds := strings.Fields(cmd)
	cmdl.textrecv.Print("command", cmd)

	if len(cmds) > 1 {
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
						_, err := cmdl.ui.CreateConversation(cmds[2], cmds[3])
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
							path := cmdl.ui.ChooseNewPath(length)
							cmdl.textrecv.Print("new path", strings.Join(path, " "))
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
				receiver := cmds[1]
				receiverServer := cmds[2]
				message := strings.Join(cmds[3:], " ")
				err := cmdl.ui.SendTo(message, receiver, receiverServer)
				if err != nil {
					cmdl.textrecv.Print("Error: ", err.Error())
				} else {
					cmdl.textrecv.Print("Message sent to ", cmds[1]+" "+cmds[2])
				}
			} else {
				cmdl.textrecv.Print("error", "usage: send receiver receiverServer message")
			}
			break
		case "exit":
			cmdl.run = false
			break
		}
	}
}