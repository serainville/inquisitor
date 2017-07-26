package plugins

import (
	"github.com/shirou/gopsutil/mem"
	"strconv"
)

func GetCPU() string {
	vmStat, _ := mem.VirtualMemory()
	return strconv.FormatUint(vmStat.Total, 10)
}
