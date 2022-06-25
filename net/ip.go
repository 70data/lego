package net

import (
	"net"
	"strings"

	"k8s.io/klog/v2"
)

func GetIP() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	localIP := strings.Split(addrs[1].String(), "/")[0]
	klog.Infoln("local ip:", localIP)
}
