package plugins

import (
	"github.com/shirou/gopsutil/mem"
	"strconv"
)

// GetCPU outputs the system's CPU usage
func GetCPU() string {
	vmStat, _ := mem.VirtualMemory()
	return strconv.FormatUint(vmStat.Total, 10)
}
