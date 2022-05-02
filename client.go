package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var addr = "localhost:8080"

func main() {
	conn, err := net.Dial("tcp", addr)
	checkForErr(err)

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewReader(conn)

	for {
		inputFromClient, err := in.ReadBytes(byte('\n'))
		checkForErr(err)

		conn.Write(inputFromClient)

		replyFromServer, err := out.ReadBytes(byte('\n'))
		checkForErr(err)

		fmt.Print("Hello, " + string(replyFromServer))
	}
}

func checkForErr(err error) {
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
}