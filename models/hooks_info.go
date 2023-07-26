package models

import (
	"encoding/json"
	"strings"
)

type HookInfoDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Args        string `json:"args"`
	ReturnValue string `json:"return-value"`
}

func addStyleToHeader(header string) string {
	return MakeTextBold(strings.ToUpper(header))
}

func (h *HookInfoDto) ToString() string {
	res := addStyleToHeader("Name:") + "                " + MakeTextHighlight(h.Name) + "\n"
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

func MakeHookInfoRequest() *Request {
	req := &Request{
		Type:    "change-info",
		Payload: EmptyPayload{},
	}
	return req
}

func hookInfoPayloadFormatter(payload any) (string, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", CliError{Message: "can`t process response payload", Cause: err}
	}

	var infoPayload HookInfoResponse
	err = json.Unmarshal(payloadBytes, &infoPayload)
	if err != nil {
		return "", CliError{Message: "can`t process response payload", Cause: err}
	}

	return infoPayload.ToString(), nil
}

func HooksInfoResponseHandler() ResponseHandler {
	return &DefaultResponseHandler{PayloadFormatter: hookInfoPayloadFormatter}
}
