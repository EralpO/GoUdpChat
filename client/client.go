package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10001")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer Conn.Close()
	for {
		buffer := make([]byte, 1024)
		println("Please enter a message")
		os.Stdin.Read(buffer)
		_, err := Conn.Write(buffer)
		if err != nil {
			fmt.Println(buffer, err)
		}
		time.Sleep(time.Second * 1)
	}
}
