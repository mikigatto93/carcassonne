package entity

import "encoding/json"

type RequestEvent struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type ResponseEvent struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

/*
type MessagePayload struct {
	Message string `json:"message"`
	From    string `json:"from"`
}*/

//type EventHandler func(event Event, client *Client) error
