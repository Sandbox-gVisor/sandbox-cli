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

type PayloadFormatter func(payload any) string

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

	var payloadText string
	if handler.PayloadFormatter == nil {
		payloadText = fmt.Sprintf("%v", response.Payload)
	} else {
		payloadText = handler.PayloadFormatter(response.Payload)
	}

	output := fmt.Sprintf("Type:      %s;\nMessage:   %s\nPayload:   %v\n",
		responseType, response.Message, payloadText)

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
