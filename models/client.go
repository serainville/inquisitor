package models

type Client struct {
	accountId		uint64	`json:"accountid"`
	Hostname		string	`json:"hostname"`
	OS				string	`json:"os"`
	ClientVersion	string	`json:"clientversion"`
}