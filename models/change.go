package models

import (
	"fmt"
)

type CallbackDto struct {
	Sysno          int    `json:"sysno"`
	EntryPoint     string `json:"entry-point"`
	CallbackSource string `json:"source"`
	Type           string `json:"type"`
}

type ChangeRequest struct {
	Type      string        `json:"type"`
	Callbacks []CallbackDto `json:"callbacks"`
}

type ChangeResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (r *ChangeResponse) ToString() string {
	return fmt.Sprintf("Type:      %s;\nMessage:   %s", r.Type, r.Message)
}
