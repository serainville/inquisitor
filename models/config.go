package models

// Describes the server configuration
type ServerConfig struct {
	IP          string
	Port        string
	TLSKeyFile  string
	TLSCertFile string
	UseTLS      bool
	Daemon      bool
	Standalone  bool
}

// Describes the client configuration
type ClientConfig struct {
	Server      string
	Port        string
	TLSKeyFile  string
	TLSCertFile string
	UseTLS      bool
}
