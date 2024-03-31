package commands

import (
	"sandbox-cli/internal/communication"
)

type DeleteHookJson struct {
	Sysno int    `json:"sysno"`
	Type  string `json:"type"`
}

type DeleteRequest struct {
	Options string           `json:"options"`
	List    []DeleteHookJson `json:"list"`
}

func MakeDeleteHooksRequest(options string, sysno int, hookType string) *communication.Request {
	var list []DeleteHookJson
	if options != "all" {
		list = append(list, DeleteHookJson{
			Type:  hookType,
			Sysno: sysno,
		})
	}
	payload := DeleteRequest{
		Options: options,
		List:    list,
	}

	req := &communication.Request{
		Type:    "unregister-hooks",
		Payload: payload,
	}
	return req
}
