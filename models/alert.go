package models

import (
	"time"
	"errors"
)

// Alert describes a generated alert message
type Alert struct {
	Name		uint64		`json:"name"`
	Message		string		`json:"message"`
	Severity	string		`json:"severity"`
	Timestamp	time.Time	`json:"timestamp"`
}

func (a *Alert) postAlert() error {
	return errors.New("Not implemeneted")
}