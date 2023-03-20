package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	target := "scanme.nmap.org"

	for port := 1; port <= 23; port++ {
		address := fmt.Sprintf("%s:%d", target, port)
		conn, err := net.DialTimeout("tcp", address, time.Second)

		if err != nil {
			fmt.Printf("Port %d is closed\n", port)
			continue
		}

		defer conn.Close()
		fmt.Printf("Port %d is open\n", port)
	}
}
