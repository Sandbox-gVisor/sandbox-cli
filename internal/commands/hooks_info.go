package commands

import (
	"encoding/json"
	"sandbox-cli/internal/communication"
	"sandbox-cli/internal/errors"
	po "sandbox-cli/internal/pretty_output"
	"strings"
)

type HookInfoDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Args        string `json:"args"`
	ReturnValue string `json:"return-value"`
}

func addStyleToHeader(header string) string {
	return po.MakeTextBold(strings.ToUpper(header))
}

func (h *HookInfoDto) ToString() string {
	res := addStyleToHeader("Name:") + "                " + po.MakeTextHighlight(h.Name) + "\n"
	res += addStyleToHeader("Description") + "          " + h.Description + "\n"
	res += addStyleToHeader("Args") + "                 " + h.Args + "\n"
	res += addStyleToHeader("Return values") + "        " + h.ReturnValue + "\n"
	return res
}

type HookInfoResponse struct {
	Hooks []HookInfoDto `json:"hooks"`
}

func (r *HookInfoResponse) ToString() string {
	res := "\n\n"
	for _, hook := range r.Hooks {
		res += hook.ToString() + "\n\n"
	}
	return res
}

func MakeHookInfoRequest() *communication.Request {
	req := &communication.Request{
		Type:    "hooks-info",
		Payload: communication.EmptyPayload{},
	}
	return req
}

func hookInfoPayloadFormatter(payload any) (string, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", errors.CliError{Message: "can`t process response payload", Cause: err}
	}

	var infoPayload HookInfoResponse
	err = json.Unmarshal(payloadBytes, &infoPayload)
	if err != nil {
		return "", errors.CliError{Message: "can`t process response payload", Cause: err}
	}

	return infoPayload.ToString(), nil
}

func HooksInfoResponseHandler() po.ResponseFormatter {
	return &po.DefaultResponseFormatter{PayloadFormatter: hookInfoPayloadFormatter}
}
