package commands

import (
	"encoding/json"
	"sandbox-cli/internal/communication"
	"sandbox-cli/internal/errors"
	po "sandbox-cli/internal/pretty_output"
	"strings"
)

type AccessorInfoDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Args        string `json:"args"`
	ReturnValue string `json:"return-value"`
}

func addStyleToHeader(header string) string {
	return po.MakeTextBold(strings.ToUpper(header))
}

func (d *AccessorInfoDto) ToString() string {
	res := addStyleToHeader("Name:") + "                " + po.MakeTextHighlight(d.Name) + "\n"
	res += addStyleToHeader("Description") + "          " + d.Description + "\n"
	res += addStyleToHeader("Args") + "                 " + d.Args + "\n"
	res += addStyleToHeader("Return values") + "        " + d.ReturnValue + "\n"
	return res
}

type AccessorInfoResponse struct {
	Accessors []AccessorInfoDto `json:"accessors"`
}

func (r *AccessorInfoResponse) ToString() string {
	res := "\n\n"
	for _, accessor := range r.Accessors {
		res += accessor.ToString() + "\n\n"
	}
	return res
}

func MakeAccessorsInfoRequest() *communication.Request {
	req := &communication.Request{
		Type:    "accessors-info",
		Payload: communication.EmptyPayload{},
	}
	return req
}

func accessorInfoPayloadFormatter(payload any) (string, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", errors.CliError{Message: "can`t process response payload", Cause: err}
	}

	var infoPayload AccessorInfoResponse
	err = json.Unmarshal(payloadBytes, &infoPayload)
	if err != nil {
		return "", errors.CliError{Message: "can`t process response payload", Cause: err}
	}

	return infoPayload.ToString(), nil
}

func AccessorInfoResponseHandler() po.ResponseFormatter {
	return &po.DefaultResponseFormatter{PayloadFormatter: accessorInfoPayloadFormatter}
}
