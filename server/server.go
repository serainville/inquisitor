package server

import (
	"fmt"
	"net/http"
	"io"
	"log"
	"os"
)

type Config struct {
	ip string
	port string
	TLSCertFile string
	TLSKeyFile string
	UseTLS bool
	Database
}
type Database struct {
	host string
	port string
	user string
	password string
}

func Init(
	bindIp string,
	bindPort string,
	TLSCertFile string,
	TLSKeyFile string,
	UseTLS bool) *Config {

	return &Config{ip: bindIp, port: bindPort, TLSCertFile: TLSCertFile, TLSKeyFile: TLSKeyFile, UseTLS: UseTLS}

}

func StartServer(c *Config) bool {

	http.HandleFunc("/", serve_root)
	http.HandleFunc("/publish", serve_publish)

	if !c.UseTLS { 
		fmt.Println("Starting server...")
		fmt.Println("")
		fmt.Println("WARNING: Server is not secure!")
		fmt.Println("         TLS should be enabled to encrypt communications.")
		fmt.Println("         Do not run in production environment!")
		fmt.Println("")
		fmt.Printf("Listening on http://%s:%s\n\n", c.ip, c.port)
		http.ListenAndServe(c.ip + ":" + c.port, nil)
	} else {
		fmt.Println("Starting server...")
		if checkTLSCert(c.TLSCertFile) && checkTLSKey(c.TLSKeyFile) {
			fmt.Printf("Listening on https://%s:%s\n\n", c.ip, c.port)
			log.Fatal(http.ListenAndServeTLS(":" + c.port, c.TLSKeyFile, c.TLSCertFile, nil))
		} else {
			log.Fatal("Could not start server")
		}
	}
	return true
}

func checkTLSCert(cert_file string) bool {
	if len(cert_file) > 0 {
		if _, err := os.Stat(cert_file); !os.IsNotExist(err) {
			return true
		}
		log.Fatal("Certificate file not found - " + cert_file)
		return false
	}
	log.Fatal("A certificate file must be specified!")
	return false
}

func checkTLSKey(key_file string) bool {
	if len(key_file) > 0 {
		if _, err := os.Stat(key_file); !os.IsNotExist(err) {
			return true
		}
		log.Fatal("Certificate key file not found - " + key_file)
		return false
	}
	log.Fatal("A certificate key file must be specified!")
	return false
}

func serve_root(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func serve_publish(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Publishing the metrics!")
}