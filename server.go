package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var addr = "localhost:8080"

func customPrint(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadBytes(byte('\n'))
		checkForErr(err)

		conn.Write(line)
	}
}

func main() {
	listener, err := net.Listen("tcp", addr)
	checkForErr(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("ERROR", err)
			continue
		}
		go customPrint(conn)
	}
}

func checkForErr(err error) {
	if(err != nil) {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
}