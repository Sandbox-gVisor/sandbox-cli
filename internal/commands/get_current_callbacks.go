package commands

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sandbox-cli/internal/communication"
	"sandbox-cli/internal/errors"
	"sandbox-cli/internal/prettyoutput"
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
			color:    prettyoutput.GreenColorText,
		},
		{
			patterns: []string{`\(`, `\)`},
			color:    prettyoutput.OrangeColorText,
		},
		{
			patterns: []string{`"`, "{", "}"},
			color:    prettyoutput.RedColorText,
		},
	}

	removeEscaping := func(str string) string {
		return strings.Replace(str, `\`, "", -1)
	}

	for _, scheme := range schemes {
		for _, pattern := range scheme.patterns {
			reg := regexp.MustCompile(pattern)
			jsSource = reg.ReplaceAllString(jsSource, prettyoutput.MakeTextBoldAndColored(removeEscaping(pattern), scheme.color))
		}
	}

	return jsSource
}

func (cj *CallbackJson) ToString(isVerbose bool) string {
	res := fmt.Sprintf("Type:          %s\n", prettyoutput.MakeTextBoldAndColored(cj.Type, prettyoutput.OrangeColorText))
	res += fmt.Sprintf("Sysno:         %s\n", prettyoutput.MakeTextBoldAndColored(strconv.Itoa(cj.Sysno), prettyoutput.OrangeColorText))
	res += fmt.Sprintf("Entry-point:   %s\n", prettyoutput.MakeTextBoldAndColored(cj.EntryPoint, prettyoutput.OrangeColorText))
	strArgs := fmt.Sprintf("%v", strings.Join(cj.CallbackArgs, ", "))
	res += fmt.Sprintf("Args:          %s\n", prettyoutput.MakeTextBoldAndColored(strArgs, prettyoutput.OrangeColorText))
	if isVerbose {
		res += fmt.Sprintf("Body:\n\n%s", highlightJsSyntax(cj.CallbackBody))
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

func MakeGetCallbacksPayloadFormatter(isVerbose bool) prettyoutput.PayloadFormatter {
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

func GetCallbackResponseHandler(isVerbose bool) prettyoutput.ResponseFormatter {
	return &prettyoutput.DefaultResponseFormatter{PayloadFormatter: MakeGetCallbacksPayloadFormatter(isVerbose)}
}

func (r *GetCallbacksPayload) ToString(isVerbose bool) string {
	res := "\n\nCallbacks:\n\n"
	for _, c := range r.Callbacks {
		res += c.ToString(isVerbose)
	}
	return res
}
