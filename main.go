package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

const (
	ServerAddr    = "localhost"
	ServerPort    = 12345
	ServerNetwork = "tcp"
)

func main() {
	l, err := net.Listen(ServerNetwork, fmt.Sprintf("%s:%d", ServerAddr, ServerPort))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	io.Copy(conn, conn)
}
