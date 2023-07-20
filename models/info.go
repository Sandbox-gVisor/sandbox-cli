package models

type InfoRequest struct {
	Type string `json:"type"`
}

type Hook struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Args        string `json:"args"`
	ReturnValue string `json:"return_value"`
}

func (h *Hook) ToString() string {
	res := "Name:         " + h.Name + "\n"
	res += "Description   " + h.Description + "\n"
	res += "Args          " + h.Args + "\n"
	res += "Return values " + h.ReturnValue + "\n"
	return res
}

type InfoResponse struct {
	Type  string `json:"type"`
	Hooks []Hook `json:"hooks"`
}

func (r *InfoResponse) ToString() string {
	res := "Type:   %s" + r.Type + "\nhooks:\n"
	for _, hook := range r.Hooks {
		res += hook.ToString()
	}
	return res
}
