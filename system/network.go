package system

import (
	"net"
)

func GetNICInfo(name string) (map[string]string, error) {
	res := make(map[string]string)
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return res, err
	}
	for _, netInterface := range netInterfaces {
		if netInterface.Name == name {
			addrs, _ := netInterface.Addrs()
			res["nic"] = name
			res["ip"] = addrs[0].String()
		}
	}
	return res, nil
}
