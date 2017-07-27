package models

// ServerConfig describes the server configuration
type ServerConfig struct {
	IP          string
	Port        string
	TLSKeyFile  string
	TLSCertFile string
	UseTLS      bool
	Daemon      bool
	Standalone  bool
}

// ClientConfig describes the client configurationd
type ClientConfig struct {
	Server      string
	Port        string
	TLSKeyFile  string
	TLSCertFile string
	UseTLS      bool
}
