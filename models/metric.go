package models

import (
	"time"
	"errors"
)

// Metric describes a measurement from a system or application
type Metric struct {
	ClientID	uint64		`json:"clientid"`
	Name		string		`json:"name"`
	Value		string		`json:"value"`
	Timestamp	time.Time	`json:"timestamp"`
}

func (m *Metric) postMetric() error {
	return errors.New("Not implemeneted")
}