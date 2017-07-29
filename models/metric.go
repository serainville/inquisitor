package models

import (
	//"errors"
	"time"
)

// MetricGroup describes a group of metrics. Used to group together related metrics.
type MetricGroup struct {
	Name    string    `json:"name"`
	Metrics []*Metric `json:"metrics"`
}

// Metric describes a measurement from a system or application
type Metric struct {
	Name  string `json:"name"`
	Group string `json:"group"`
	Value string `json:"value"`
}

// ClientMetrics describes a Client's metrics message
type ClientMetrics struct {
	AccountID string         `json:"accountid"`
	HostID    string         `json:"hostid"`
	Hostname  string         `json:"hostname"`
	Timestamp time.Time      `json:"timestamp"`
	Secret    string         `json:"secret"`
	Groups    []*MetricGroup `'json:"groups"`
	OS        string         `'json:"os"`
	Platform  string         `'json:"platform"`
	Kernel    string         `'json:"kernel"`
}

/*
func (m *Metric) postMetric() error {
	return errors.New("Not implemeneted")
}
*/
