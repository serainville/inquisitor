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

func sendMetrics(client *models.ClientMetrics, host, port string) (bool, error) {
	js, _ := json.Marshal(client)
	url := "http://" + host + ":" + port + "/api/v1/metrics"
	consoleLog.Info("Publishing to: " + url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(js))
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

	return true, nil
}

func collectMetrics() (m *models.ClientMetrics) {
	client := &models.ClientMetrics{
		ClientID:  1010101,
		Secret:    "a44ecab3784ad4545",
		Timestamp: time.Now(),
	}

	Metrica := []*models.Metric{}

	Metrica = append(Metrica, appendMetric("total", "memory", plugins.GetMemoryTotal()))
	Metrica = append(Metrica, appendMetric("free", "memory", plugins.GetMemoryFree()))
	Metrica = append(Metrica, appendMetric("used", "memory", plugins.GetMemoryUsed()))
	//Metrica = append(Metrica, appendMetric("total", "memory", plugins.GetNumberRunningProcess()))

	client.Metrics = Metrica

	return client
}

func appendMetric(name string, group string, value string) *models.Metric {
	m := new(models.Metric)
	m.Name = name
	m.Group = group
	m.Value = value
	return m
}
