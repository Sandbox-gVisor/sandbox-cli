package models

import "fmt"

type GetRequest struct {
	Type string `json:"type"`
}
type CallbackJson struct {
	Type       string `json:"type"`
	Sysno      int    `json:"sysno"`
	EntryPoint string `json:"entry-point"`
	Src        string `json:"source"`
}

func (cj *CallbackJson) ToString() string {
	res := fmt.Sprintf("  Type:          %s\n", cj.Type)
	res += fmt.Sprintf("  Sysno:         %d\n", cj.Sysno)
	res += fmt.Sprintf("  Entry-point:   %s\n", cj.EntryPoint)
	res += fmt.Sprintf("  Src:           %s\n\n\n", cj.EntryPoint)
	return res
}

type GetResponse struct {
	Type      string         `json:"type"`
	Callbacks []CallbackJson `json:"callbacks"`
	Message   string         `json:"message"`
}

func (r *GetResponse) ToString() string {
	res := fmt.Sprintf("Type:           %s\n", r.Type)
	res += fmt.Sprintf("Message:        %s\n", r.Message)
	res += "Callbacks:\n"
	for _, c := range r.Callbacks {
		res += c.ToString()
	}
	return res
}
