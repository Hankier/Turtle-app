package cmdsListener

import (
	"client/textReceiver"
	"bufio"
	"os"
	"strings"
	"strconv"
	"client/client/ui"
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

func (cmdl *Listener)execGet(cmds []string){
	cmdnum := len(cmds)

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
}

func (cmdl *Listener)execNew(cmds []string) {
	cmdnum := len(cmds)

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
}
func (cmdl *Listener)execSet(cmds []string) {
	cmdnum := len(cmds)

	if cmdnum > 1 {
		switch cmds[1] {
		case "pathenc":
			if cmdnum > 2 {
				enctype := cmds[2]
				switch enctype {
				case "PLAIN":
					cmdl.ui.SetEncryptionType(crypt.PLAIN)
					cmdl.textrecv.Print("set pathenc", "PLAIN")
					break
				case "RSA":
					cmdl.ui.SetEncryptionType(crypt.RSA)
					cmdl.textrecv.Print("set pathenc", "RSA")
					break
				case "ELGAMAL":
					cmdl.ui.SetEncryptionType(crypt.ELGAMAL)
					cmdl.textrecv.Print("set pathenc", "ELGAMAL")
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
				convokeyfn := cmds[4]
				enctype := cmds[5]
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
				cmdl.textrecv.Print("error", "usage: set convokey receiverServer receiver filename [RSA|ELGAMAL]")
			}
		}
	} else {
		cmdl.textrecv.Print("error", "usage: set [pathenc|convokey]")
	}
}

func (cmdl *Listener)execSend(cmds []string) {
	cmdnum := len(cmds)

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
}

func (cmdl *Listener)execConnect(cmds []string) {
	cmdnum := len(cmds)

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
		//todo remove debug
		case "d":
			cmdl.execCmd("connect server01")
			cmdl.execCmd("new convo server01 client01")
			cmdl.execCmd("set pathenc ELGAMAL")
			cmdl.execCmd("set convokey server01 client01 publicKeyElGamal ELGAMAL")
			cmdl.execCmd("send server01 client01 a")
			break
		case "get":
			cmdl.execGet(cmds)
			break
		case "connect":
			cmdl.execConnect(cmds)
			break
		case "new":
			cmdl.execNew(cmds)
			break
		case "set":
			cmdl.execSet(cmds)
			break
		case "send":
			cmdl.execSend(cmds)
			break
		case "exit":
			cmdl.run = false
			break
		case "help":
			cmdl.textrecv.Print("help", "\n" +
				"get\n" +
				"    path - shows current path\n" +
				"    servers - shows available server\n" +
				"\n" +
				"set\n" +
				"    pathenc\n" +
				"        PLAIN - sets path encryption to plain text\n" +
				"        RSA - sets path encryption to RSA\n" +
				"        ELGAMAL - sets path encryption to ElGamal\n" +
				"    convokey\n" +
				"        receiverServer receiver filename\n" +
				"            RSA - sets RSA public key for convo\n" +
				"            ELGAMAL - sets ElGamal public key for convo\n" +
				"\n" +
				"connect\n" +
				"    serverName - connects to given server\n" +
				"\n" +
				"new\n" +
				"    convo\n" +
				"        receiverServer receiver - creates new conversation with given names\n" +
				"    path\n" +
				"        length - creates new random path with given length\n" +
				"\n" +
				"send\n" +
				"    receiverServer receiver message - sends message with given data" +
				"\n")
			break
		default:
			cmdl.textrecv.Print("error", "avaiable cmds: get set connect new send help exit")
		}
	}
}