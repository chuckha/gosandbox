package main

import (
	"fmt"
	"net"
)

func main () {
	ln, err := net.Listen("unix", "/tmp/go.sock")
	if err != nil { panic(err) }
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil { fmt.Println(err) }
		fmt.Println("We got a connection!")
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	buf := make([]byte, 1024)
	for {
		numBytes, err := c.Read(buf)
		if err != nil { panic(err) }

		data := buf[:numBytes]
		fmt.Println(string(data))
	}
}
