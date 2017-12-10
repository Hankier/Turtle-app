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
	userInterface userInterface.UserInterface
	textReceiver textReceiver.TextReceiver
	run bool
}

func NewCommandsListener(userInterface userInterface.UserInterface, textReceiver textReceiver.TextReceiver)(*CommandsListener){
	cmdl := new(CommandsListener)
	cmdl.userInterface = userInterface
	cmdl.textReceiver = textReceiver
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
	cmdl.textReceiver.Print("command", cmd)

	if len(cmds) > 1 {
		switch cmds[0] {
		case "get":
			if len(cmds) > 1 {
				switch cmds[1] {
				case "path":
					cmdl.textReceiver.Print("path", strings.Join(cmdl.userInterface.GetCurrentPath(), " "))
				case "servers":
					cmdl.textReceiver.Print("servers", strings.Join(cmdl.userInterface.GetServerList(), " "))
				}
			} else {
				cmdl.textReceiver.Print("error", "usage: get path, get servers")
			}
			break
		case "connect":
			if len(cmds) > 1 {
				err := cmdl.userInterface.ConnectToServer(cmds[1])
				if err != nil {
					cmdl.textReceiver.Print("Error: ", "Wrong server")
				} else {
					cmdl.textReceiver.Print("Connecting to server ", cmds[1])
				}
			} else {
				cmdl.textReceiver.Print("error", "usage: connect serverName")
			}
			break
		case "new":
			if len(cmds) > 1 {
				switch cmds[1] {
				case "convo":
					if len(cmds) > 3 {
						_, err := cmdl.userInterface.CreateConversation(cmds[2], cmds[3])
						if err != nil {
							cmdl.textReceiver.Print("Error: ", err.Error())
						} else {
							cmdl.textReceiver.Print("Created conversation ", cmds[2]+" "+cmds[3])
						}
					} else {
						cmdl.textReceiver.Print("error", "usage: new convo clientName serverName")
					}
				case "path":
					if len(cmds) > 2 {
						length, err := strconv.Atoi(cmds[2])
						if err != nil {
							cmdl.textReceiver.Print("Error: ", err.Error())
						} else {
							path := cmdl.userInterface.ChooseNewPath(length)
							cmdl.textReceiver.Print("new path", strings.Join(path, " "))
						}
					} else {
						cmdl.textReceiver.Print("error", "usage: new path length")
					}
				}
			} else {
				cmdl.textReceiver.Print("error", "usage: new convo clientName serverName, new path length")
			}
			break
		case "send":
			if len(cmds) > 2 {
				receiver := cmds[1]
				receiverServer := cmds[2]
				message := strings.Join(cmds[3:], " ")
				err := cmdl.userInterface.SendTo(message, receiver, receiverServer)
				if err != nil {
					cmdl.textReceiver.Print("Error: ", err.Error())
				} else {
					cmdl.textReceiver.Print("Message sent to ", cmds[1]+" "+cmds[2])
				}
			} else {
				cmdl.textReceiver.Print("error", "usage: send receiver receiverServer message")
			}
			break
		case "exit":
			cmdl.run = false
			break
		}
	}
}