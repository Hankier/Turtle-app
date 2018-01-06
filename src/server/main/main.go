package main

import (
	"os"
	"server/server"
	"log"
)

func main() {
	args := os.Args[1:]

	name := args[0]
	clientsPort := args[1]
	serversPort := args[2]
	log.Println("My name is " + name)
	srv := server.NewServer(name)
	srv.Start(clientsPort, serversPort)

}