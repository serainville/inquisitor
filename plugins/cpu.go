package plugins

import (
	"github.com/shirou/gopsutil/cpu"
	"strconv"
)

// GetCPU outputs the system's CPU usage

func GetCPUIdle() string {
	percentage, _ := cpu.Percent(0, false)

	cpuc := 0.0

	for idx, cpua := range percentage {
		cpuc = cpuc + cpua
		idx = idx + idx
	}

	return strconv.FormatInt(int64(cpuc), 10)

}
