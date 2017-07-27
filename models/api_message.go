package models

// Message describes the json-formatted output message of an API request
type Message struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
