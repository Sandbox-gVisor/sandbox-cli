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

type HookJson struct {
	Sysno      int      `json:"sysno"`
	EntryPoint string   `json:"entry-point"`
	Source     string   `json:"source"`
	Body       string   `json:"body"`
	Args       []string `json:"args"`
	Type       string   `json:"type"`
}

func (hj *HookJson) ToString(isVerbose bool) string {
	res := fmt.Sprintf("Type:          %s\n", po.MakeTextBoldAndColored(hj.Type, po.OrangeColorText))
	res += fmt.Sprintf("Sysno:         %s\n", po.MakeTextBoldAndColored(strconv.Itoa(hj.Sysno), po.OrangeColorText))
	res += fmt.Sprintf("Entry-point:   %s\n", po.MakeTextBoldAndColored(hj.EntryPoint, po.OrangeColorText))
	strArgs := fmt.Sprintf("%v", strings.Join(hj.Args, ", "))
	res += fmt.Sprintf("Args:          %s\n", po.MakeTextBoldAndColored(strArgs, po.OrangeColorText))
	if isVerbose {
		res += fmt.Sprintf("Body:\n\n%s", po.HighlightJsSyntax(hj.Body))
	}
	res += "\n\n"
	return res
}

type GetHooksPayload struct {
	Hooks []HookJson `json:"hooks"`
}

func MakeGetHooksRequest() *communication.Request {
	req := &communication.Request{
		Type:    "current-hooks",
		Payload: communication.EmptyPayload{},
	}
	return req
}

func MakeGetHooksPayloadFormatter(isVerbose bool) po.PayloadFormatter {
	return func(payload any) (string, error) {
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return "", errors.CliError{Message: "can`t process response payload", Cause: err}
		}

		var getHooksPayload GetHooksPayload
		err = json.Unmarshal(payloadBytes, &getHooksPayload)
		if err != nil {
			return "", errors.CliError{Message: "can`t process response payload", Cause: err}
		}

		return getHooksPayload.ToString(isVerbose), nil
	}
}

func GetHookResponseHandler(isVerbose bool) po.ResponseFormatter {
	return &po.DefaultResponseFormatter{PayloadFormatter: MakeGetHooksPayloadFormatter(isVerbose)}
}

func (r *GetHooksPayload) ToString(isVerbose bool) string {
	res := "\n\nHooks:\n\n"
	for _, c := range r.Hooks {
		res += c.ToString(isVerbose)
	}
	return res
}
