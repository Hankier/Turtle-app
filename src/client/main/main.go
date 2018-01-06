package main

import (
	"os"
	"client/client"
)

func main(){
	args := os.Args[1:]

	name := args[0]

	cli := client.New(name)
	if args[1] != "" {
		cli.ConnectToServer(args[1])
	}
	cli.Start()
}