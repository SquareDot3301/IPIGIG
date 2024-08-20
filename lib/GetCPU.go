package lib

import (
	"github.com/shirou/gopsutil/v3/cpu"
)

func GetCPU() (string, error) {
	info, err := cpu.Info()
	if err != nil {
		return "", err
	}

	if len(info) > 0 {
		return info[0].ModelName, nil
	}

	return "", nil
}
