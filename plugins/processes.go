package plugins

import (
	"github.com/shirou/gopsutil/host"
	"strconv"
)

func GetNumberRunningProcess() string {
	hostStat, _ := host.Info()
	return strconv.FormatUint(hostStat.Procs, 10)
}
