package models

type CallbackDto struct {
	Sysno          int    `json:"sysno"`
	EntryPoint     string `json:"entry-point"`
	CallbackSource string `json:"source"`
	Type           string `json:"type"`
}

type ChangeRequest struct {
	Callbacks []CallbackDto `json:"callbacks"`
}

func MakeChangeCallbacksRequest(dtos []CallbackDto) *Request {
	req := &Request{
		Type:    "change-callbacks",
		Payload: ChangeRequest{Callbacks: dtos},
	}
	return req
}
