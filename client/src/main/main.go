package main

import (
	"os"
	"client"
)

func main(){
	args := os.Args[1:]

	name := args[0]

	cli := client.New(name)
	cli.Start()
}