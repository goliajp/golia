package libutil

import (
	"net"
	"strings"
)

func GetMacAddr() (string, error) {
	addrArr, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	var currentIP, currentNetworkHardwareName string
	for _, address := range addrArr {
		if IPNet, ok := address.(*net.IPNet); ok && !IPNet.IP.IsLoopback() {
			if IPNet.IP.To4() != nil {
				currentIP = IPNet.IP.String()
				break
			}
		}
	}

	interfaces, _ := net.Interfaces()

Exit:
	for _, netInterface := range interfaces {
		if addrArr, err := netInterface.Addrs(); err == nil {
			for _, addr := range addrArr {
				if strings.Contains(addr.String(), currentIP) {
					currentNetworkHardwareName = netInterface.Name
					break Exit
				}
			}
		}
	}

	netInterface, err := net.InterfaceByName(currentNetworkHardwareName)
	if err != nil {
		return "", err
	}

	macAddr := netInterface.HardwareAddr.String()
	return macAddr, nil
}
