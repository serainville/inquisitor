package plugins

import (
	"github.com/shirou/gopsutil/host"
	"strconv"
)

// GetNumberRunningProcess returns the number of running processes
func GetNumberRunningProcess() string {
	hostStat, _ := host.Info()
	return strconv.FormatUint(hostStat.Procs, 10)
}
