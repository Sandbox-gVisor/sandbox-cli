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

type CallbackDto struct {
	Sysno          int    `json:"sysno"`
	EntryPoint     string `json:"entry-point"`
	CallbackSource string `json:"source"`
}

type Request struct {
	Type string `json:"type"`

	Callbacks []CallbackDto `json:"callbacks"`
}

type Response struct {
	Type string `json:"type"`
}
