package main

import (
	"net"
	"fmt"
)


func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "beaglebone.local:7527")
	if err != nil {
		fmt.Println("Could not create address:", err)
	}

	con, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Could not connect to UDP addresss:", err)
	}
	defer con.Close()

	buf := []byte("hello")
	_, err = con.Write(buf)
	if err != nil {
		fmt.Println("Could not write to connection:", err)
	}


}
