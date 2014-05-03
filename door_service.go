package main

import (
	"encoding/json"
	"fmt"
	"net"
)

var (
	localPort = 6666
)

func sendMsg(socket *net.UDPConn, msg string) {
	// MarshalIndent makes sure our JSON is pretty
	jmsg, err := json.MarshalIndent(msg, "", "  ")
	if err != nil {
		fmt.Println("Could not marshal message:", err)
	}

	// But MarshalIndent doesn't come with a newline, so we do that ourselves.
	jmsg = append(jmsg, "\n"...)

}

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
