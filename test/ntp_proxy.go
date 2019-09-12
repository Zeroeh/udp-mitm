package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

var (
	udpListener  *net.UDPConn
	udpRemoteConn *net.UDPConn
	laddr         *net.UDPAddr
	raddr         *net.UDPAddr
	counts        = 0
	gamePort      = ":123"
	gameIP        = "time.apple.com" //use wireshark to grab this from the game server
)

func main() {
	network, err := net.ResolveUDPAddr("udp", ":5555")
	if err != nil {
		fmt.Println("Error resolving:", err)
	}
	udpListener, err = net.ListenUDP("udp", network)
	if err != nil {
		fmt.Println("Error starting listener:", err)
	}
	raddr, err = net.ResolveUDPAddr("udp", gameIP+gamePort)
	if err != nil {
		fmt.Println("Error resolving remote:", err)
	}
	go localReader()
	for {
		time.Sleep(60 * time.Second)
	}
}

func localReader() {
	for {
		var err error
		var bRead int
		var addr *net.UDPAddr
		buf := make([]byte, 8096)
		bRead, addr, err = udpListener.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Err reading local:", err)
		}
		if strings.Contains(addr.String(), gamePort) { //forward to local
			fmt.Println("S2C -> ", buf[:bRead])
			_, err = udpListener.WriteToUDP(buf[:bRead], laddr)
			if err != nil {
				fmt.Println("Err forwarding to local:", err)
			}
		} else { //forward to remote
			laddr = addr //reset so we send to the correct port
			fmt.Println("C2S -> ", buf[:bRead])
			_, err = udpListener.WriteToUDP(buf[:bRead], raddr)
			if err != nil {
				fmt.Println("Err forwarding to remote:", err)
			}
		}
	}
}

