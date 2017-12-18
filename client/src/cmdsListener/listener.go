package cmdsListener

import (
	"textReceiver"
	"bufio"
	"os"
	"strings"
	"strconv"
	"client/ui"
	"log"
	"crypt"
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
		if len(text) > 0 {
			text = text[:len(text) - 1] //cut newline
		}
		if err != nil{
			log.Println(err)
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

	cmdnum := len(cmds)

	if cmdnum > 0 {
		switch cmds[0] {
		case "get":
			if cmdnum > 1 {
				switch cmds[1] {
				case "path":
					cmdl.textrecv.Print("path", strings.Join(cmdl.ui.GetCurrentPath(), " "))
				case "servers":
					if cmdnum > 2{
						if cmds[2] == "details"{
							details := "\n"
							serverList := cmdl.ui.GetServerList()
							for _, server := range serverList {
								srvDetails := cmdl.ui.GetServerDetails(server)
								details += server + "\n"
								for _, srvDetail := range srvDetails{
									details += srvDetail + "\n"
								}
								details += "\n"
							}
							cmdl.textrecv.Print("servers details", details)
						} else {
							cmdl.textrecv.Print("error", "usage: get servers details")
						}

					} else {
						cmdl.textrecv.Print("servers", strings.Join(cmdl.ui.GetServerList(), " "))
					}
				}
			} else {
				cmdl.textrecv.Print("error", "usage: get path, get servers")
			}
			break
		case "connect":
			if cmdnum > 1 {
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
			if cmdnum > 1 {
				switch cmds[1] {
				case "convo":
					if cmdnum > 3 {
						err := cmdl.ui.CreateConversation(cmds[2], cmds[3])
						if err != nil {
							cmdl.textrecv.Print("Error: ", err.Error())
						} else {
							cmdl.textrecv.Print("Created conversation ", cmds[2]+" "+cmds[3])
						}
					} else {
						cmdl.textrecv.Print("error", "usage: new convo receiverServer receiver")
					}
				case "path":
					if cmdnum > 2 {
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
				cmdl.textrecv.Print("error", "usage: new convo receiverServer receiver, new path length")
			}
			break
		case "set":
			if cmdnum > 1 {
				switch cmds[1] {
				case "pathenc":
					if cmdnum > 2 {
						enctype := cmds[2]
						switch enctype {
						case "PLAIN":
							cmdl.ui.SetEncryptionType(crypt.PLAIN)
							break
						case "RSA":
							cmdl.ui.SetEncryptionType(crypt.RSA)
							break
						case "ELGAMAL":
							cmdl.ui.SetEncryptionType(crypt.ELGAMAL)
							break
						default:
							cmdl.textrecv.Print("error", "use [PLAIN|RSA|ELGAMAL]")
						}
					} else {
						cmdl.textrecv.Print("error", "usage: set pathenc [PLAIN|RSA|ELGAMAL]")
					}
					break
				case "convokey":
					if cmdnum > 5 {
						recServer := cmds[2]
						rec := cmds[3]
						enctype := cmds[4]
						convokeyfn := cmds[5]
						switch enctype {
						case "RSA":
							err := cmdl.ui.SetConversationKey(recServer, rec, crypt.RSA, convokeyfn)
							if err != nil{
								cmdl.textrecv.Print("error", err.Error())
							} else {
								cmdl.textrecv.Print("set convokey", "RSA succesful")
							}
							break
						case "ELGAMAL":
							err := cmdl.ui.SetConversationKey(recServer, rec, crypt.ELGAMAL, convokeyfn)
							if err != nil{
								cmdl.textrecv.Print("error", err.Error())
							} else {
								cmdl.textrecv.Print("set convokey", "ELGAMAL succesful")
							}
							break
						default:
							cmdl.textrecv.Print("error", "use [RSA|ELGAMAL]")
							break
						}
					} else {
						cmdl.textrecv.Print("error", "usage: set convokey receiverServer receiver[RSA|ELGAMAL] filename")
					}
				}
			} else {
				cmdl.textrecv.Print("error", "usage: set [pathenc|convokey]")
			}
		case "send":
			if cmdnum > 2 {
				receiverServer := cmds[1]
				receiver := cmds[2]
				message := strings.Join(cmds[3:], " ")
				err := cmdl.ui.SendTo(receiverServer, receiver, message)
				if err != nil {
					cmdl.textrecv.Print("error", err.Error())
				} else {
					cmdl.textrecv.Print("sent", cmds[1]+" "+cmds[2])
				}
			} else {
				cmdl.textrecv.Print("error", "usage: send receiverServer receiver message")
			}
			break
		case "exit":
			cmdl.run = false
			break
		case "help":
			if cmdnum > 1 {
				switch cmds[2]{
				case "get":
					cmdl.textrecv.Print("help", "usage: get path, get servers")
					break
				case "set":
					cmdl.textrecv.Print("help", "usage: set [pathenc|convokey]")
					break
				case "connect":
					cmdl.textrecv.Print("help", "usage: connect serverName")
					break
				case "new":
					cmdl.textrecv.Print("help", "usage: new [convo|path]")
					break
				case "send":
					cmdl.textrecv.Print("help", "usage: send receiverServer receiver message")
					break
				}
			} else {
				cmdl.textrecv.Print("help", "avaiable cmds: get set connect new send exit")
			}
			break
		default:
			cmdl.textrecv.Print("error", "avaiable cmds: get set connect new send exit")
		}
	}
}