package models

import (
	//"errors"
	"time"
)

// Metric describes a measurement from a system or application
type Metric struct {
	Name  string
	Group string
	Value string
}

// ClientMetrics describes a Client's metrics message
type ClientMetrics struct {
	ClientID  uint64
	Timestamp time.Time
	Secret    string
	Metrics   []*Metric
}

/*
func (m *Metric) postMetric() error {
	return errors.New("Not implemeneted")
}
*/
