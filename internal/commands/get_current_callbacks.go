package commands

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sandbox-cli/internal/communication"
	"sandbox-cli/internal/errors"
	"sandbox-cli/internal/pretty_output"
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

func highlightJsSyntax(jsSource string) string {
	type colorScheme struct {
		patterns []string
		color    int
	}

	schemes := []colorScheme{
		{
			patterns: []string{"function", "if", "return", "const", "let", "var"},
			color:    pretty_output.GreenColorText,
		},
		{
			patterns: []string{`\(`, `\)`},
			color:    pretty_output.OrangeColorText,
		},
		{
			patterns: []string{`"`, "{", "}"},
			color:    pretty_output.RedColorText,
		},
	}

	removeEscaping := func(str string) string {
		return strings.Replace(str, `\`, "", -1)
	}

	for _, scheme := range schemes {
		for _, pattern := range scheme.patterns {
			reg := regexp.MustCompile(pattern)
			jsSource = reg.ReplaceAllString(jsSource, pretty_output.MakeTextBoldAndColored(removeEscaping(pattern), scheme.color))
		}
	}

	return jsSource
}

func (cj *CallbackJson) ToString(isVerbose bool) string {
	res := fmt.Sprintf("Type:          %s\n", pretty_output.MakeTextBoldAndColored(cj.Type, pretty_output.OrangeColorText))
	res += fmt.Sprintf("Sysno:         %s\n", pretty_output.MakeTextBoldAndColored(strconv.Itoa(cj.Sysno), pretty_output.OrangeColorText))
	res += fmt.Sprintf("Entry-point:   %s\n", pretty_output.MakeTextBoldAndColored(cj.EntryPoint, pretty_output.OrangeColorText))
	strArgs := fmt.Sprintf("%v", cj.CallbackArgs)
	res += fmt.Sprintf("Args:          %s\n", pretty_output.MakeTextBoldAndColored(strArgs, pretty_output.OrangeColorText))
	if isVerbose {
		res += fmt.Sprintf("Body:\n\n%s", highlightJsSyntax(cj.CallbackBody))
	}
	res += "\n\n\n"
	return res
}

type GetCallbacksPayload struct {
	Callbacks []CallbackJson `json:"callbacks"`
}

func MakeGetCallbacksRequest() *communication.Request {
	req := &communication.Request{
		Type:    "current-callbacks",
		Payload: pretty_output.EmptyPayload{},
	}
	return req
}

func MakeGetCallbacksPayloadFormatter(isVerbose bool) pretty_output.PayloadFormatter {
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

func GetCallbackResponseHandler(isVerbose bool) pretty_output.ResponseHandler {
	return &pretty_output.DefaultResponseHandler{PayloadFormatter: MakeGetCallbacksPayloadFormatter(isVerbose)}
}

func (r *GetCallbacksPayload) ToString(isVerbose bool) string {
	res := "\n\nCallbacks:\n"
	for _, c := range r.Callbacks {
		res += c.ToString(isVerbose)
	}
	return res
}
