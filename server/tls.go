package server

import (
	"log"
	"os"
)

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