package models

import (
	//"errors"
	"time"
)

// MetricGroup describes a group of metrics. Used to group together related metrics.
type MetricGroup struct {
	Name string `json:"name"`
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
	ClientID  uint64    `json:"clientid"`
	Timestamp time.Time `json:"timestamp"`
	Secret    string    `json:"secret"`
	Groups	  []*MetricGroup `'json:"groups"`
	//Metrics   []*Metric `json:"metrics"`
}



/*
func (m *Metric) postMetric() error {
	return errors.New("Not implemeneted")
}
*/
