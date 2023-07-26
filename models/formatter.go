package models

import "fmt"

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
