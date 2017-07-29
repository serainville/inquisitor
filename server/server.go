package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	//"github.com/nytimes/gziphandler"
	"github.com/serainville/gologger"

	"github.com/serainville/inquisitor/models"
	"github.com/serainville/inquisitor/plugins"
	"github.com/serainville/inquisitor/storage/influxdb"
	"github.com/serainville/inquisitor/variables"
)

var consoleLog = gologger.GetLogger(gologger.BASIC, gologger.ColoredLog)

// StartStandalone starts Inquisitors standalone server
func StartStandalone() {
	consoleLog.Info("Starting in Standalone mode")
	consoleLog.Warn("This feature is not fully implemented")
	consoleLog.Info("Initilizing agents [cpu, memory, network, storage")

	for {
		consoleLog.Info("Inquiring...")

		go fmt.Println("Mem Total: " + plugins.GetMemoryTotal())
		go fmt.Println("Mem Free: " + plugins.GetMemoryFree())
		go fmt.Println("Mem Used: " + plugins.GetMemoryUsed())
		go fmt.Println("# Process: " + plugins.GetNumberRunningProcess())
		go fmt.Println(plugins.GetNetwork())
		go fmt.Println(plugins.GetStorage())

		time.Sleep(30 * time.Second)
		// running standlone. Ctrl-C to kill
	}
	// Implement client agent runs
	// Implement Monitor storage
	// Implement APM storage
}

// StartServer starts the Inquisitor HTTP server
func StartServer(c *models.ServerConfig) bool {

	logger1 := gologger.GetLogger(gologger.BASIC, gologger.ColoredLog)

	http.HandleFunc("/", serveRoot)
	http.HandleFunc("/api/v1/metrics", receiveMetrics)
	http.HandleFunc("/api/v1/apm", receiveAPM)

	if !c.UseTLS {
		logger1.Info("Starting server...")
		logger1.Info("Version: " + variables.Version + " (" + variables.CommitID + ")")
		logger1.Warn("WARNING: TLS disabled. Server is not secure!")
		logger1.Warn("WARNING: Do not use in production")
		logger1.Info("Listening on http://" + c.IP + ":" + c.Port)

		log.Fatal(http.ListenAndServe(c.IP+":"+c.Port, nil))
	} else {
		logger1.Info("Starting server...")
		fmt.Println("Starting server...")
		if checkTLSCert(c.TLSCertFile) && checkTLSKey(c.TLSKeyFile) {
			fmt.Printf("Listening on https://%s:%s\n\n", c.IP, c.Port)
			log.Fatal(http.ListenAndServeTLS(c.IP+":"+c.Port, c.TLSKeyFile, c.TLSCertFile, nil))
		} else {
			log.Fatal("Could not start server")
		}
	}
	return true
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "API V1 Running")
}

func receiveMetrics(w http.ResponseWriter, r *http.Request) {
	message := models.Message{200, "Metric saved successfully"}

	data, err := ioutil.ReadAll(r.Body)
	consoleLog.Info(string(data))

	var t models.ClientMetrics

	json.Unmarshal(data, &t)
	if err != nil {
		panic(err)
	}

	err = influxdb.WriteMetrics(t)
	if err != nil {
		consoleLog.Error("Metrics write failed.")
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}

	js, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func receiveAPM(w http.ResponseWriter, r *http.Request) {

}

func appendMetric(name string, group string, value string) *models.Metric {
	m := new(models.Metric)
	m.Name = name
	m.Group = group
	m.Value = value
	return m
}
