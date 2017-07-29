package plugins

import (
	"github.com/shirou/gopsutil/mem"
	"strconv"
)

// GetMemoryTotal returns a system's total memory
func GetMemoryTotal() string {
	vmStat, _ := mem.VirtualMemory()
	return strconv.FormatUint(vmStat.Total, 10)
}

// GetMemoryFree returns a system's free memory
func GetMemoryFree() string {
	vmStat, _ := mem.VirtualMemory()
	return strconv.FormatUint(vmStat.Free, 10)
}

// GetMemoryUsed outputs a systems total used memory
func GetMemoryUsed() string {
	vmStat, _ := mem.VirtualMemory()
	return strconv.FormatUint(vmStat.Used, 10)
}

func GetMemoryAvailable() string {
	vmStat, _ := mem.VirtualMemory()
	return strconv.FormatUint(vmStat.Available, 10)
}

func GetMemoryUsedPercent() string {
	vmStat, _ := mem.VirtualMemory()
	return strconv.FormatInt(int64(vmStat.UsedPercent), 10)
}
