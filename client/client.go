package client

import (
	"bytes"
	"encoding/json"
	//"fmt"
	"net/http"
	"time"
	//"io/ioutil"

	"github.com/serainville/gologger"
	"github.com/serainville/inquisitor/models"
	"github.com/serainville/inquisitor/plugins"
	"github.com/shirou/gopsutil/host"
)

var consoleLog = gologger.GetLogger(gologger.BASIC, gologger.ColoredLog)

type Client struct {
	Hostname string
	IP       string
	OS       string
	Platform string
	Kernel   string
	CPUCores uint16
	RAM      uint64
}

func (c Client) register() error {
	// Store client info in database
	return nil
}

func (c Client) getClientInfo() error {
	hostStat, _ := host.Info()

	c.Hostname = hostStat.Hostname
	c.IP = "192.168.1.1"
	c.OS = hostStat.OS
	c.Platform = hostStat.Platform
	c.Kernel = ""
	c.CPUCores = 1
	c.RAM = 32000000

	return nil
}

func (c Client) unregister() error {
	return nil
}

func (c Client) collect(host string, port string, pl plugins.Plugins) error {

	for {
		client := collectMetrics()
		sendMetrics(client, host, port)
		time.Sleep(5 * time.Second)
	}
	return nil
}

func (c Client) init() (plugins.Plugins) {

	c.getClientInfo()
	c.checkRegistration()
	pl := c.loadPlugins()
	
	


	return pl
}

func (c Client) loadPlugins() plugins.Plugins {
	pl := plugins.Plugins{}
	pl.Discover()
	pl.List()
	return pl
}

func (c Client) isRegistered() (bool, error) {
	return true, nil
}

func (c Client) checkRegistration() error {
	consoleLog.Info("Registering server....")
	isRegistered, err := c.isRegistered()
	if err != nil {
		return err
	}

	if isRegistered != true {
		err := c.register()
		if err != nil {
			return err
		}
	}
	consoleLog.Info("Server already registered")
	return nil
}


// Start starts the client service
func Start(host, port string) {
	consoleLog.Info("Starting Inquisitor client")

	c	:= new(Client)

	//consoleLog.Info("Connecting to server " + host + ":" + port + "....")

	pl := c.init()

	c.collect(host, port, pl)

}

func sendMetrics(client *models.ClientMetrics, host, port string) {
	js, _ := json.Marshal(client)
	url := "http://" + host + ":" + port + "/api/v1/metrics"
	consoleLog.Info("Publishing to: " + url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(js))
	if err != nil {

	}
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		consoleLog.Error("Could not communicate with server.")
	} else {

		consoleLog.Info(resp.Status)
		defer resp.Body.Close()
	}
}

func collectMetrics() (m *models.ClientMetrics) {

	hostStat, _ := host.Info()
	hostname := hostStat.Hostname

	client := &models.ClientMetrics{
		AccountID: "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		HostID:    hostStat.HostID,
		OS:        hostStat.OS,
		Platform:  hostStat.Platform,
		Kernel:    hostStat.KernelVersion,
		Hostname:  hostname,
		Secret:    "a44ecab3784ad4545",
		Timestamp: time.Now(),
	}

	MemoryMetrics := []*models.Metric{}
	MemoryMetrics = append(MemoryMetrics, appendMetric("total", "memory", plugins.GetMemoryTotal()))
	MemoryMetrics = append(MemoryMetrics, appendMetric("free", "memory", plugins.GetMemoryFree()))
	MemoryMetrics = append(MemoryMetrics, appendMetric("used", "memory", plugins.GetMemoryUsed()))
	MemoryMetrics = append(MemoryMetrics, appendMetric("available", "memory", plugins.GetMemoryAvailable()))
	MemoryMetrics = append(MemoryMetrics, appendMetric("percent", "memory", plugins.GetMemoryUsedPercent()))
	//Metrica = append(Metrica, appendMetric("total", "memory", plugins.GetNumberRunningProcess()))

	cpuMetrics := []*models.Metric{}
	cpuMetrics = append(cpuMetrics, appendMetric("percent", "cpu", plugins.GetCPUIdle()))

	hostMetrics := []*models.Metric{}
	hostMetrics = append(hostMetrics, appendMetric("uptime", "host", plugins.GetUptime()))

	groups := []*models.MetricGroup{}
	groups = append(groups, appendMetricGroup("memory", MemoryMetrics))
	groups = append(groups, appendMetricGroup("cpu", cpuMetrics))
	groups = append(groups, appendMetricGroup("host", hostMetrics))

	client.Groups = groups

	return client
}

func appendMetricGroup(name string, metrics []*models.Metric) *models.MetricGroup {
	g := new(models.MetricGroup)
	g.Name = name
	g.Metrics = metrics
	return g
}

func appendMetric(name string, group string, value string) *models.Metric {
	m := new(models.Metric)
	m.Name = name
	m.Group = group
	m.Value = value
	return m
}
