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

	peers := make(chan net.Conn, 1)
	go pair(peers)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		peers <- conn
	}
}

func pair(peers chan net.Conn) {
	for {
		p1, p2 := <-peers, <-peers
		go chat(p1, p2)
	}
}

func chat(p1, p2 net.Conn) {
	go transfer(p1, p2)
	go transfer(p2, p1)
}

func transfer(dst, src net.Conn) {
	io.Copy(dst, src)
	dst.Close()
}
