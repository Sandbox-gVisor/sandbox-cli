package models

import (
	"strings"
)

type Hook struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Args        string `json:"args"`
	ReturnValue string `json:"return-value"`
}

func addStyleToHeader(header string) string {
	return MakeTextBold(strings.ToUpper(header))
}

func (h *Hook) ToString() string {
	res := addStyleToHeader("Name:") + "                " + MakeTextHighlight(h.Name) + "\n"
	res += addStyleToHeader("Description") + "          " + h.Description + "\n"
	res += addStyleToHeader("Args") + "                 " + h.Args + "\n"
	res += addStyleToHeader("Return values") + "        " + h.ReturnValue + "\n"
	return res
}

type InfoResponse struct {
	Hooks []Hook `json:"hooks"`
}

func MakeHookInfoRequest() *Request {
	req := &Request{
		Type:    "change-info",
		Payload: EmptyPayload{},
	}
	return req
}
