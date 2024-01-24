package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	serverConn, err := net.ListenUDP("udp", &net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: 10001})
	if err != nil {
		log.Fatalln(err)
	}
	defer serverConn.Close()
	for {
		buf := make([]byte, 1024)
		n, _, err := serverConn.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Result: ", string(buf[:n]))
	}
}
