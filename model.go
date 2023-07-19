package main

import (
	"encoding/json"
)

type Message struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

func (m *Message) ToString() string {
	b, err := json.Marshal(m)
	if err != nil {
		return "{}"
	}
	return string(b)
}
