package models

type StateRequest struct {
	EntryPoint string `json:"entry-point"`
	Src        string `json:"source"`
}

type StateResponse struct {
	Type string `json:"type"`
}

func (r *StateResponse) ToString() string {
	return "Type:   " + r.Type
}

func MakeChangeStateRequest(fileName string) *Request {
	payload := StateRequest{
		Src:        string(ReadFile(fileName)),
		EntryPoint: "",
	}

	req := &Request{
		Type:    "change-state",
		Payload: payload,
	}
	return req
}
