package main

import (
	"fmt"
	"net"
)

var (
	localPort = 6666
)

func main() {
	buff := make([]byte, 1024)

	// bind to local port
	port := fmt.Sprintf(":%d", localPort)
	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		fmt.Println("Error when creating local address:", err)
	}

	// create udp socket
	sock, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error when creating socket:", err)
	}

	fmt.Println("Listening to port:", localPort)

	for {
		// we do not need the length, remote Addr is going to be the BeagleBones address
		rlen, remoteAddr, err := sock.ReadFromUDP(buff)
		if err != nil {
			fmt.Println("Error reading from socket:", err)
		}

		fmt.Println("Received packets from:", remoteAddr)

		fmt.Println(string(buff[0:rlen]))
	}

}
