package models

import (
	"errors"
)

// Client describes a server registered to Inquisitor
type Client struct {
	AccountID		uint64	`json:"accountid"`
	Hostname		string	`json:"hostname"`
	OS				string	`json:"os"`
	ClientVersion	string	`json:"clientversion"`
}

func (m *Client) postClient() error {
	return errors.New("Not implemeneted")
}