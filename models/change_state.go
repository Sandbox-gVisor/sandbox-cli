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

func MakeChangeStateRequest(entryPoint string, src string) *Request {
	payload := StateRequest{
		Src:        src,
		EntryPoint: entryPoint,
	}

	req := &Request{
		Type:    "change-state",
		Payload: payload,
	}
	return req
}
