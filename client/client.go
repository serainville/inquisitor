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
)

var consoleLog = gologger.GetLogger(gologger.BASIC, gologger.ColoredLog)

// Start starts the client service
func Start(host, port string) {

	consoleLog.Info("Starting Inquisitor client")
	consoleLog.Warn("This feature is not fully implemented")
	consoleLog.Info("Connecting to server " + host + ":" + port + "....")
	consoleLog.Info("Initilizing agents [cpu, memory, network, storage]")

	for {
		consoleLog.Info("Gathing metrics...")

		client := collectMetrics()

		sendMetrics(client, host, port)

		time.Sleep(5 * time.Second)
	}
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
		// fmt.Printf("%s\n", js)
		//htmlData, _ := ioutil.ReadAll(resp.Body)
		consoleLog.Info(resp.Status)
		//fmt.Printf("%s\n%s\n", resp.Status)
		defer resp.Body.Close()
	}
}

func collectMetrics() (m *models.ClientMetrics) {
	client := &models.ClientMetrics{
		ClientID:  1010101,
		Secret:    "a44ecab3784ad4545",
		Timestamp: time.Now(),
	}



	MemoryMetrics := []*models.Metric{}
	MemoryMetrics = append(MemoryMetrics, appendMetric("total", "memory", plugins.GetMemoryTotal()))
	MemoryMetrics = append(MemoryMetrics, appendMetric("free", "memory", plugins.GetMemoryFree()))
	MemoryMetrics = append(MemoryMetrics, appendMetric("used", "memory", plugins.GetMemoryUsed()))
	//Metrica = append(Metrica, appendMetric("total", "memory", plugins.GetNumberRunningProcess()))

	groups := []*models.MetricGroup{}
	groups = append(groups, appendMetricGroup("memory", MemoryMetrics))

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
