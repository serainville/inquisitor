package models

import (
	"time"
)

type Metric struct {
	ClientId	uint64		`json:"clientid"`
	Name		string		`json:"name"`
	Value		string		`json:"value`
	Timestamp	time.Time	`json:"timestamp"`
}

func (m *Metric) postMetric() error {
	return errors.New("Not implemeneted")
}