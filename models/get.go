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

func (cj *CallbackJson) ToString(isVerbose bool) string {
	res := fmt.Sprintf("  Type:          %s\n", cj.Type)
	res += fmt.Sprintf("  Sysno:         %d\n", cj.Sysno)
	res += fmt.Sprintf("  Entry-point:   %s\n", cj.EntryPoint)
	if isVerbose {
		res += fmt.Sprintf("  Src:           %s", cj.Src)
	}
	res += "\n\n\n"
	return res
}

type GetResponse struct {
	Type      string         `json:"type"`
	Callbacks []CallbackJson `json:"callbacks"`
	Message   string         `json:"message"`
}

func (r *GetResponse) ToString(isVerbose bool) string {
	res := fmt.Sprintf("Type:           %s\n", r.Type)
	res += fmt.Sprintf("Message:        %s\n", r.Message)
	res += "Callbacks:\n"
	for _, c := range r.Callbacks {
		res += c.ToString(isVerbose)
	}
	return res
}
