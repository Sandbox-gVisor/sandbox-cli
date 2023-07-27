package models

import (
	"encoding/json"
	"fmt"
	"regexp"
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
			color:    GreenColorText,
		},
		{
			patterns: []string{`\(`, `\)`},
			color:    OrangeColorText,
		},
		{
			patterns: []string{`"`, "{", "}"},
			color:    RedColorText,
		},
	}

	removeEscaping := func(str string) string {
		return strings.Replace(str, `\`, "", -1)
	}

	for _, scheme := range schemes {
		for _, pattern := range scheme.patterns {
			reg := regexp.MustCompile(pattern)
			jsSource = reg.ReplaceAllString(jsSource, MakeTextBoldAndColored(removeEscaping(pattern), scheme.color))
		}
	}

	return jsSource
}

func (cj *CallbackJson) ToString(isVerbose bool) string {
	res := fmt.Sprintf("Type:          %s\n", MakeTextBoldAndColored(cj.Type, OrangeColorText))
	res += fmt.Sprintf("Sysno:         %s\n", MakeTextBoldAndColored(strconv.Itoa(cj.Sysno), OrangeColorText))
	res += fmt.Sprintf("Entry-point:   %s\n", MakeTextBoldAndColored(cj.EntryPoint, OrangeColorText))
	strArgs := fmt.Sprintf("%v", cj.CallbackArgs)
	res += fmt.Sprintf("Args:          %s\n", MakeTextBoldAndColored(strArgs, OrangeColorText))
	if isVerbose {
		res += fmt.Sprintf("Body:\n\n%s", highlightJsSyntax(cj.CallbackBody))
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
