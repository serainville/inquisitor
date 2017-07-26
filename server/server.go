package server

import (
	"fmt"
	"net/http"
	"io"
	"log"
	"os"
	"encoding/json"
	"time"

	//"github.com/nytimes/gziphandler"
	"github.com/serainville/gologger"

	"github.com/serainville/inquisitor/variables"
	"github.com/serainville/inquisitor/models"
)

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
		http.ListenAndServe(c.IP + ":" + c.Port, nil)
	} else {
		logger1.Info("Starting server...")
		fmt.Println("Starting server...")
		if checkTLSCert(c.TLSCertFile) && checkTLSKey(c.TLSKeyFile) {
			fmt.Printf("Listening on https://%s:%s\n\n", c.IP, c.Port)
			log.Fatal(http.ListenAndServeTLS(c.IP + ":" + c.Port, c.TLSKeyFile, c.TLSCertFile, nil))
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
	message := models.Message{200,"Metric saved successfully"}

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
	m := models.Metric{101010101,"cpu_load","45",time.Now()}

	js, err := json.Marshal(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

}