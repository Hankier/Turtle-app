package main

import (
	"fmt"
	"net"
	"bufio"
	"strconv"
)

var conns = make([]net.Conn, 0, 2)

func sendToConns(src net.Conn, msg string){
	for i := 0; i < len(conns); i++{
		if conns[i] != src {
			conns[i].Write([]byte(msg))
		}
	}
}

func main(){
	fmt.Println("Launching server...")

	ln, _ := net.Listen("tcp",":8081")

	fmt.Println("Listening...")

	for {
		conn, _ := ln.Accept()
		fmt.Println("Client accepted " + conn.LocalAddr().String())
		conns = append(conns, conn)
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for{
		msg, err := reader.ReadString('\n')
		if err != nil{
			fmt.Println(err)
			break
		}
		fmt.Println("Got message:" + msg + " len:" + strconv.Itoa(len(msg)))
		sendToConns(conn, msg)
	}
	conn.Close()
}
