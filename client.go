package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
)

func run() {
	conn, _ := net.Dial("tcp", "localhost:8081")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text + "\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
