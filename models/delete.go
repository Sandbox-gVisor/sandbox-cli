package models

type DeleteCallbackJson struct {
	Sysno int    `json:"sysno"`
	Type  string `json:"type"`
}

type DeleteRequest struct {
	Type    string               `json:"type"`
	Options string               `json:"options"`
	List    []DeleteCallbackJson `json:"list"`
}

type DeleteResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (dr *DeleteResponse) ToString() string {
	return "Type:    " + dr.Type + "\nMessage:  " + dr.Message + "\n"
}
