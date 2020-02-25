package net

import (
	"log"
	"net"
	"strings"
)

// LocalIP is make local ip
var LocalIP string

// GetIP is get local ip
func GetIP() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	LocalIP = strings.Split(addrs[1].String(), "/")[0]
	log.Println("local ip:", LocalIP)
}
