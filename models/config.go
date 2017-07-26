package models

type ServerConfig struct {
	IP			string
	Port		string
	TLSKeyFile	string
	TLSCertFile	string
	UseTLS		bool
	Daemon		bool
}

type ClientConfig struct {
	Server		string
	Port		string
	TLSKeyFile	string
	TLSCertFile	string
	UseTLS		bool
}