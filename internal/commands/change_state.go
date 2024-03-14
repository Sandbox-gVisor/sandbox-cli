package commands

import (
	"sandbox-cli/internal/communication"
)

type StateRequest struct {
	Src string `json:"source"`
}

func MakeChangeStateRequest(fileName string) *communication.Request {
	payload := StateRequest{
		Src: string(ReadFile(fileName)),
	}

	req := &communication.Request{
		Type:    "change-state",
		Payload: payload,
	}
	return req
}
