package client

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/serainville/gologger"
	"github.com/serainville/inquisitor/models"
	"github.com/serainville/inquisitor/plugins"
)

var consoleLog = gologger.GetLogger(gologger.BASIC, gologger.ColoredLog)

// Start starts the client service
func Start() {

	consoleLog.Info("Starting Inquisitor client")
	consoleLog.Warn("This feature is not fully implemented")
	consoleLog.Info("Connecting to server....")
	consoleLog.Info("Initilizing agents [cpu, memory, network, storage]")

	for {
		consoleLog.Info("Inquiring...")

		client := &models.ClientMetrics{
			ClientID:  1010101,
			Secret:    "a44ecab3784ad4545",
			Timestamp: time.Now(),
		}

		Metrica := []*models.Metric{}

		Metrica = append(Metrica, appendMetric("total", "memory", plugins.GetMemoryTotal()))
		Metrica = append(Metrica, appendMetric("free", "memory", plugins.GetMemoryFree()))
		Metrica = append(Metrica, appendMetric("used", "memory", plugins.GetMemoryUsed()))
		Metrica = append(Metrica, appendMetric("total", "memory", plugins.GetNumberRunningProcess()))		

		client.Metrics = Metrica
		// Convert struct to json
		js, _ := json.Marshal(client)
	
		fmt.Printf("%s\n", js)

		time.Sleep(5 * time.Second)
	}
}



func appendMetric(name string, group string, value string) *models.Metric {
	m := new(models.Metric)
	m.Name = name
	m.Group = group
	m.Value = value
	return m
}