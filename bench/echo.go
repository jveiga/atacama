package main

import (
	"bufio"
	"fmt"
	"io"
	"fmt"
	"net"
	"os"
)

func main() {
	// Listen on TCP port 2112 on all interfaces.
	l, err := net.Listen("tcp", ":2112")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on :2112")

	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	const buffSize = 1024*50
	// Create a buffer to hold the received data.
	r := bufio.NewReaderSize(conn, buffSize)

	for {
		_, err := io.Copy(conn, r)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading:", err.Error())
			}
			break
		}
	}
}
