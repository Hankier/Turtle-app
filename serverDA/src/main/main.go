package main

import (
	"os"
	"server"
)

func main() {
	args := os.Args[1:]

	name := args[0]
	clientsPort := args[1]
	serversPort := args[2]

	srv := server.NewServer(name)
	srv.Start(clientsPort, serversPort)

}
