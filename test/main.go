package main

import (
	"fmt"
	"net"
)

var (
	udpConn *net.UDPConn
	counts = 0
)

func main() {
	network, err := net.ResolveUDPAddr("udp", ":5555")
	if err != nil {
		fmt.Println("Error resolving:", err)
	}
	udpConn, err = net.ListenUDP("udp", network)
	if err != nil {
		fmt.Println("Error starting listener:", err)
	}
	for {
		buf := make([]byte, 1024)
		read, addr, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading:", err)
		}
		fmt.Println("Got packet from", addr)
		counts++
		fmt.Printf("Packet: %d | Buffer: %v\n", counts, buf[:read])
		fmt.Println("Attempting to forward request...")
		ntpHost, err := net.ResolveUDPAddr("udp", "time.apple.com:123")
		fmt.Println("Sending to", ntpHost)
		udpConn.WriteToUDP(buf[:read], ntpHost)
		buf2 := make([]byte, 1024)
		read2, addr2, err := udpConn.ReadFromUDP(buf2)
		if err != nil {
			fmt.Println("Error reading packet:", err)
		}
		fmt.Println("Got NTP response from", addr2)
		fmt.Println("Data:", buf2[:read2])
		udpConn.WriteToUDP(buf2[:read2], addr) //send the packet back
		fmt.Println("Sent the packet back to", addr)
	}
}

