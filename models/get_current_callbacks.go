package models

import (
	"encoding/json"
	"fmt"
)

type GetRequest struct {
	Type string `json:"type"`
}
type CallbackJson struct {
	Type       string `json:"type"`
	Sysno      int    `json:"sysno"`
	EntryPoint string `json:"entry-point"`
	Src        string `json:"source"`
}

func (cj *CallbackJson) ToString(isVerbose bool) string {
	res := fmt.Sprintf("  Type:          %s\n", cj.Type)
	res += fmt.Sprintf("  Sysno:         %d\n", cj.Sysno)
	res += fmt.Sprintf("  Entry-point:   %s\n", cj.EntryPoint)
	if isVerbose {
		res += fmt.Sprintf("  Src:           %s", cj.Src)
	}
	res += "\n\n\n"
	return res
}

type GetCallbacksPayload struct {
	Callbacks []CallbackJson `json:"callbacks"`
}

func MakeGetCallbacksRequest() *Request {
	req := &Request{
		Type:    "current-callbacks",
		Payload: EmptyPayload{},
	}
	return req
}

func MakeGetCallbacksPayloadFormatter(isVerbose bool) PayloadFormatter {
	return func(payload any) (string, error) {
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return "", CliError{Message: "can`t process response payload", Cause: err}
		}

		var getCallbacksPayload GetCallbacksPayload
		err = json.Unmarshal(payloadBytes, &getCallbacksPayload)
		if err != nil {
			return "", CliError{Message: "can`t process response payload", Cause: err}
		}

		return getCallbacksPayload.ToString(isVerbose), nil
	}
}

func GetCallbackResponseHandler(isVerbose bool) ResponseHandler {
	return &DefaultResponseHandler{MakeGetCallbacksPayloadFormatter(isVerbose)}
}

func (r *GetCallbacksPayload) ToString(isVerbose bool) string {
	res := "\n\nCallbacks:\n"
	for _, c := range r.Callbacks {
		res += c.ToString(isVerbose)
	}
	return res
}
