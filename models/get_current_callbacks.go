package models

import (
	"encoding/json"
	"fmt"
	"strconv"
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
	res := fmt.Sprintf("  Type:          %s\n", MakeTextBoldAndColored(cj.Type, OrangeColorText))
	res += fmt.Sprintf("  Sysno:         %s\n", MakeTextBoldAndColored(strconv.Itoa(cj.Sysno), OrangeColorText))
	res += fmt.Sprintf("  Entry-point:   %s\n", MakeTextBoldAndColored(cj.EntryPoint, OrangeColorText))
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
