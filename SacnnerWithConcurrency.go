package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	target := "scanme.nmap.org"
	var wg sync.WaitGroup

	for port := 1; port <= 24; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", target, port)
			conn, err := net.DialTimeout("tcp", address, time.Second)

			if err != nil {
				fmt.Printf("Port %d is closed\n", port)
				return
			}

			_ = conn.Close()
			fmt.Printf("Port %d is open\n", port)
		}(port)
	}

	wg.Wait()
}
