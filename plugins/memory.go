package plugins

import (
	"github.com/shirou/gopsutil/mem"
	"strconv"
)

func GetMemoryTotal() string {
	vmStat, _ := mem.VirtualMemory()
	return strconv.FormatUint(vmStat.Total, 10)
}

func GetMemoryFree() string {
	vmStat, _ := mem.VirtualMemory()
	return strconv.FormatUint(vmStat.Free, 10)
}

func GetMemoryUsed() string {
	vmStat, _ := mem.VirtualMemory()
	return strconv.FormatUint(vmStat.Used, 10)
}
