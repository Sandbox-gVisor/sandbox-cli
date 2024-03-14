package commands

import (
	"encoding/json"
	"fmt"
	"sandbox-cli/internal/communication"
	"sandbox-cli/internal/errors"
	po "sandbox-cli/internal/pretty_output"
	"strconv"
	"strings"
)

type CallbackJson struct {
	Sysno          int      `json:"sysno"`
	EntryPoint     string   `json:"entry-point"`
	CallbackSource string   `json:"source"`
	CallbackBody   string   `json:"body"`
	CallbackArgs   []string `json:"args"`
	Type           string   `json:"type"`
}

func (cj *CallbackJson) ToString(isVerbose bool) string {
	res := fmt.Sprintf("Type:          %s\n", po.MakeTextBoldAndColored(cj.Type, po.OrangeColorText))
	res += fmt.Sprintf("Sysno:         %s\n", po.MakeTextBoldAndColored(strconv.Itoa(cj.Sysno), po.OrangeColorText))
	res += fmt.Sprintf("Entry-point:   %s\n", po.MakeTextBoldAndColored(cj.EntryPoint, po.OrangeColorText))
	strArgs := fmt.Sprintf("%v", strings.Join(cj.CallbackArgs, ", "))
	res += fmt.Sprintf("Args:          %s\n", po.MakeTextBoldAndColored(strArgs, po.OrangeColorText))
	if isVerbose {
		res += fmt.Sprintf("Body:\n\n%s", po.HighlightJsSyntax(cj.CallbackBody))
	}
	res += "\n\n"
	return res
}

type GetCallbacksPayload struct {
	Callbacks []CallbackJson `json:"callbacks"`
}

func MakeGetCallbacksRequest() *communication.Request {
	req := &communication.Request{
		Type:    "current-callbacks",
		Payload: communication.EmptyPayload{},
	}
	return req
}

func MakeGetCallbacksPayloadFormatter(isVerbose bool) po.PayloadFormatter {
	return func(payload any) (string, error) {
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return "", errors.CliError{Message: "can`t process response payload", Cause: err}
		}

		var getCallbacksPayload GetCallbacksPayload
		err = json.Unmarshal(payloadBytes, &getCallbacksPayload)
		if err != nil {
			return "", errors.CliError{Message: "can`t process response payload", Cause: err}
		}

		return getCallbacksPayload.ToString(isVerbose), nil
	}
}

func GetCallbackResponseHandler(isVerbose bool) po.ResponseFormatter {
	return &po.DefaultResponseFormatter{PayloadFormatter: MakeGetCallbacksPayloadFormatter(isVerbose)}
}

func (r *GetCallbacksPayload) ToString(isVerbose bool) string {
	res := "\n\nCallbacks:\n\n"
	for _, c := range r.Callbacks {
		res += c.ToString(isVerbose)
	}
	return res
}
