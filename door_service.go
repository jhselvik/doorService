package main

import (
	"net"
	"fmt"
)








func main() {
	buff := make([]byte, 1024)

	// bind to local port
	addr, err := net.ResolveUDPAddr("udp", ":6666")
	if err != nil {
		fmt.Println("Error when creating local address:", err)
	}

	// create udp socket
	sock, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error when creating socket:", err)
	}

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


// CLient code
// serverAddr, err := net.ResolveUDPAddr("udp", "beaglebone.local:7527")
// if err != nil {
// 	fmt.Println("Could not create address:", err)
// }

// con, err := net.DialUDP("udp", nil, serverAddr)
// if err != nil {
// 	fmt.Println("Could not connect to UDP addresss:", err)
// }
// defer con.Close()

// buf := []byte("hello")
// _, err = con.Write(buf)
// if err != nil {
// 	fmt.Println("Could not write to connection:", err)
// }

