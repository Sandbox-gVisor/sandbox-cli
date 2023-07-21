package models

type StateRequest struct {
	Type       string `json:"type"`
	EntryPoint string `json:"entry-point"`
	Src        string `json:"source"`
}

type StateResponse struct {
	Type string `json:"type"`
}

func (r *StateResponse) ToString() string {
	return "Type:   %s" + r.Type
}
