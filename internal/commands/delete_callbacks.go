package commands

import (
	"sandbox-cli/internal/communication"
)

type DeleteCallbackJson struct {
	Sysno int    `json:"sysno"`
	Type  string `json:"type"`
}

type DeleteRequest struct {
	Options string               `json:"options"`
	List    []DeleteCallbackJson `json:"list"`
}

func MakeDeleteCallbacksRequest(options string, sysno int, callbackType string) *communication.Request {
	var list []DeleteCallbackJson
	if options != "all" {
		list = append(list, DeleteCallbackJson{
			Type:  callbackType,
			Sysno: sysno,
		})
	}
	payload := DeleteRequest{
		Options: options,
		List:    list,
	}

	req := &communication.Request{
		Type:    "unregister-callbacks",
		Payload: payload,
	}
	return req
}
