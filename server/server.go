package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	"io/ioutil"

	//"github.com/nytimes/gziphandler"
	"github.com/serainville/gologger"

	"github.com/serainville/inquisitor/models"
	"github.com/serainville/inquisitor/plugins"
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

		go fmt.Println(plugins.GetCPU())
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

func checkTLSCert(certFile string) bool {
	if len(certFile) > 0 {
		if _, err := os.Stat(certFile); !os.IsNotExist(err) {
			return true
		}
		log.Fatal("Certificate file not found - " + certFile)
		return false
	}
	log.Fatal("A certificate file must be specified!")
	return false
}

func checkTLSKey(keyFile string) bool {
	if len(keyFile) > 0 {
		if _, err := os.Stat(keyFile); !os.IsNotExist(err) {
			return true
		}
		log.Fatal("Certificate key file not found - " + keyFile)
		return false
	}
	log.Fatal("A certificate key file must be specified!")
	return false
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "API V1 Running")
}


func receiveMetrics(w http.ResponseWriter, r *http.Request) {
	message := models.Message{200, "Metric saved successfully"}

	data, err := ioutil.ReadAll(r.Body)
	consoleLog.Info(string(data))

	
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
	//m := models.Metric{101010101, "cpu_load", "45", time.Now()}
	m2 := &models.ClientMetrics{
		ClientID:  1010101,
		Secret:    "a44ecab3784ad4545",
		Timestamp: time.Now(),
	}

	Metrica := []*models.Metric{}
	//metric := new(models.Metric)

	Metrica = append(Metrica, appendMetric("free", "memory", plugins.GetMemoryTotal()))
	Metrica = append(Metrica, appendMetric("used", "memory", plugins.GetMemoryFree()))
	Metrica = append(Metrica, appendMetric("total", "memory", plugins.GetNumberRunningProcess()))

	m2.Metrics = Metrica

	js, err := json.Marshal(m2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Token", "fs4sas-fsaffsadf4g-cxgfsdgdsgdsg-gdfsgdsfg")
		w.Write(js)
	}

}

func appendMetric(name string, group string, value string) *models.Metric {
	m := new(models.Metric)
	m.Name = name
	m.Group = group
	m.Value = value
	return m
}
