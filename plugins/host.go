package plugins

import (
	"github.com/shirou/gopsutil/host"
	"strconv"
)

func GetUptime() string {
	hostStat, _ := host.Info()
	uptime := strconv.FormatUint(hostStat.Uptime, 10)
	return uptime
}

func GetHostID() string {
	hostStat, _ := host.Info()
	hostid := hostStat.HostID
	return hostid
}
