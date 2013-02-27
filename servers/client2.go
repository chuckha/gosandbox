package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main () {
	conn, err := net.Dial("unix", "/tmp/go.sock")
	defer conn.Close()
	if err != nil { panic(err) }


	reader := bufio.NewReader(os.Stdin)

	for {
		line, e := reader.ReadString('\n')
		if e != nil { panic(e) }
		fmt.Printf("You said: %s", line)

		conn.Write([]byte(line))
	}
}
