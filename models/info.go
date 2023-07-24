package models

import (
	"fmt"
	"strings"
)

type InfoRequest struct {
	Type string `json:"type"`
}

type Hook struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Args        string `json:"args"`
	ReturnValue string `json:"return-value"`
}

func addStyleToHeader(header string) string {
	return makeTextBold(strings.ToUpper(header))
}

func (h *Hook) ToString() string {
	res := addStyleToHeader("Name:") + "                " + makeTextHighlight(h.Name) + "\n"
	res += addStyleToHeader("Description") + "          " + h.Description + "\n"
	res += addStyleToHeader("Args") + "                 " + h.Args + "\n"
	res += addStyleToHeader("Return values") + "        " + h.ReturnValue + "\n"
	return res
}

type InfoResponse struct {
	Type  string `json:"type"`
	Hooks []Hook `json:"hooks"`
}

func (r *InfoResponse) ToString() string {
	res := fmt.Sprintf("Type:   %s\nhooks:\n\n", r.Type)
	for _, hook := range r.Hooks {
		res += hook.ToString() + "\n\n"
	}
	return res
}
