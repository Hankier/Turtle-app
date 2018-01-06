package main

import (
	"os"
	"client/client"
)

func main(){
	args := os.Args[1:]

	name := args[0]

	cli := client.New(name)
	if len(args) > 1 && args[1] != "" {
		cli.ConnectToServer(args[1])
	}
	cli.Start()
}