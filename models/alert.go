package models

import (
	"time"
)

type Alert struct {
	Name		uint64		`json:"name"`
	Message		string		`json:"message"`
	Severity	string		`json:"Severity`
	Timestamp	time.Time	`json:"timestamp"`
}

func (a *Alert) postAlert() error {
	return errors.New("Not implemeneted")
}