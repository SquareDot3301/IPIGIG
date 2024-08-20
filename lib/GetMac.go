package lib

import (
	"bytes"
	"fmt"
	"net"
)

func GetMAC() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, i := range interfaces {
		if i.Flags&net.FlagUp != 0 && !bytes.Equal(i.HardwareAddr, nil) {
			return i.HardwareAddr.String(), nil
		}
	}

	return "", fmt.Errorf("no MAC address found")
}
