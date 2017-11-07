package main

import (
	"fmt"
	"net"
	"bufio"
	"strconv"
	"io"
)

var conns = make([]net.Conn, 0, 2)

func sendToConns(src net.Conn, data []byte){
	for i := 0; i < len(conns); i++{
		if conns[i] != src {
			writer := bufio.NewWriter(conns[i])
			fmt.Println("Send data length: ", len(data))
			writer.Write(data)
			writer.Flush()
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

func readCmd(reader *bufio.Reader) ([]byte, error){
	data := make([]byte, 0)
	cmd, err := reader.ReadBytes('\n')
	if err != nil{return nil, err}
	data = append(data, cmd...)

	fmt.Println("Got cmd:", string(cmd[0:len(cmd) -1]))

	sizeB, err := reader.ReadBytes('\n')
	if err != nil{return nil, err}
	data = append(data, sizeB...)

	size, _ := strconv.Atoi(string(sizeB[0:len(sizeB) - 1]))
	fmt.Println("Size:", size)
	for i := 0; i < size; i++{
		elSizeB, err := reader.ReadBytes('\n')
		if err != nil{return nil, err}
		data = append(data, elSizeB...)

		elSize, _ := strconv.Atoi(string(elSizeB[0:len(elSizeB) - 1]))
		fmt.Println("El size:", elSize)


		elData := make([]byte, elSize)
		_, err = io.ReadFull(reader, elData)
		if err != nil{return nil, err}
		fmt.Println("El data:", elData)
		data = append(data, elData...)
	}

	return data, nil
}

func handleConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for{
		data, err := readCmd(reader)
		if err != nil{fmt.Println(err); break}
		sendToConns(conn, data)
	}
	conn.Close()
}
