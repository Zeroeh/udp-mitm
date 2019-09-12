package main

import (
	"fmt"
	"net"
	"time"
)

const (
	serverHost = "time.apple.com"
	serverPort = "123" //standard ntp port
)

var (
	readBuffer = make([]byte, 128)
)

func main() {
	for {
		fmt.Println("Sending NTP request...")
		clientRequest := []byte("\x1b\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
		udpAddr, err := net.ResolveUDPAddr("udp", serverHost+":"+serverPort)
		if err != nil {
			fmt.Println("Error resolving udp host:", err)
		}
		socket, err := net.DialUDP("udp", nil, udpAddr)
		if err != nil {
			fmt.Println("Error dialing server:", err)
		}
		wtn, err := socket.Write(clientRequest)
		if err != nil {
			fmt.Println("Error writing message:", err)
		}
		if wtn != len(clientRequest) {
			fmt.Printf("Difference in written msg length: Written: %d | Total: %d\n", wtn, len(clientRequest))
		}
		fmt.Println("Wrote message!")
		read, _, err := socket.ReadFromUDP(readBuffer)
		if err != nil {
			fmt.Println("Error reading message:", err)
		}
		if read == 0 {
			fmt.Println("Got blank message!?")
		}
		readBuffer = readBuffer[:read]
		fmt.Println("Received buffer!", readBuffer)
		time.Sleep(time.Second * 10)
	}
}
