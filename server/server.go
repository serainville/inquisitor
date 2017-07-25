package server

import (
	"fmt"
	"net/http"
	"io"
)

type Config struct {
	bindip string
	port string
	Database
}
type Database struct {
	host string
	port string
	user string
	password string
}

func StartServer(bindIp string, bindPort string) {
	conf := Config{bindip: bindIp, port: bindPort}

	fmt.Println("Server started.")
	fmt.Printf("Listening on http://%s:%s\n\n", conf.bindip, conf.port)
	http.HandleFunc("/", serve_root)
	http.HandleFunc("/publish", serve_publish)
	http.ListenAndServe(":"+conf.port, nil)
}

func serve_root(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func serve_publish(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Publishing the metrics!")
}