package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
)

type Request struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}

const OkResponseType = "ok"
const ErrorResponseType = "error"

type Response struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Payload any    `json:"payload"`
}

type EmptyPayload struct{}

type ResponseHandler interface {
	Handle(response *Response)
}

type PayloadFormatter func(payload any) (string, error)

func DefaultPayloadFormatter(payload any) (string, error) {
	payloadText := ""
	if payload != nil {
		payloadText = fmt.Sprintf("%v", payload)
	}
	return payloadText, nil
}

type DefaultResponseHandler struct {
	PayloadFormatter PayloadFormatter
}

func (handler *DefaultResponseHandler) Handle(response *Response) {
	var responseType string
	switch response.Type {
	case OkResponseType:
		responseType = MakeTextBoldAndColored(response.Type, GreenColorText)
	case ErrorResponseType:
		responseType = MakeTextBoldAndColored(response.Type, RedColorText)
	default:
		responseType = MakeTextBold(response.Type)
	}

	type header struct {
		name  string
		value string
	}
	var headers []header

	headers = append(headers, header{"Type", responseType})
	if response.Message != "" {
		headers = append(headers, header{"Message", response.Message})
	}

	formatter := handler.PayloadFormatter
	if formatter == nil {
		formatter = DefaultPayloadFormatter
	}

	payloadText, err := formatter(response.Payload)
	if err != nil {
		payloadText = MakeTextBoldAndColored(err.Error(), RedColorText)
	}
	if payloadText != "" {
		headers = append(headers, header{"Payload", payloadText})
	}

	output := "\n"
	for _, h := range headers {
		output += fmt.Sprintf("%s: %s\n", MakeTextBold(h.name), h.value)
	}

	fmt.Println(output)
}

func writeToConn(conn net.Conn, content []byte) error {
	for len(content) > 0 {
		n, err := conn.Write(content)
		if err != nil {
			return err
		}
		content = content[n:]
	}

	return nil
}

func SendRequest(addr string, request *Request) (*Response, error) {
	conn, err := net.Dial("unix", addr)
	if err != nil {
		return nil, MakeCliError("gvisor connection error", err)
	}
	defer func(conn net.Conn) {
		if err := conn.Close(); err != nil {
			fmt.Println(err)
		}
	}(conn)

	marshaledReqeust, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	err = writeToConn(conn, marshaledReqeust)
	if err != nil {
		return nil, MakeCliError("error sending request to gvisor", err)
	}

	rawResponse, err := io.ReadAll(conn)
	if err != nil {
		return nil, MakeCliError("error reading gvisor response", err)
	}

	var response Response
	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		return nil, MakeCliError("error decoding gvisor response", err)
	}

	return &response, nil
}
