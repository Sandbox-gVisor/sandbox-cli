package communication

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"sandbox-cli/internal/errors"
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
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, errors.MakeCliError("gvisor connection error", err)
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
		return nil, errors.MakeCliError("error sending request to gvisor", err)
	}

	rawResponse, err := io.ReadAll(conn)
	if err != nil {
		return nil, errors.MakeCliError("error reading gvisor response", err)
	}

	var response Response
	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		return nil, errors.MakeCliError("error decoding gvisor response", err)
	}

	return &response, nil
}
