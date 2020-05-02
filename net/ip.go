package net

import (
	"log"
	"net"
	"strings"
)

func GetIP() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	localIP := strings.Split(addrs[1].String(), "/")[0]
	log.Println("local ip:", localIP)
}
