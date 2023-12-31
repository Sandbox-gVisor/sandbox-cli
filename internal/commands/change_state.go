package commands

import (
	"sandbox-cli/internal/communication"
)

type StateRequest struct {
	EntryPoint string `json:"entry-point"`
	Src        string `json:"source"`
}

type StateResponse struct {
	Type string `json:"type"`
}

func (r *StateResponse) ToString() string {
	return "Type:   " + r.Type
}

func MakeChangeStateRequest(fileName string) *communication.Request {
	payload := StateRequest{
		Src:        string(ReadFile(fileName)),
		EntryPoint: "",
	}

	req := &communication.Request{
		Type:    "change-state",
		Payload: payload,
	}
	return req
}
