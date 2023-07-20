package cli

import (
	"encoding/json"
	"fmt"
	models "sandbox-cli/models"
)

func sendChange(address string, dtos []models.CallbackDto) string {
	req := models.ChangeRequest{
		Type:      "change-callbacks",
		Callbacks: dtos,
	}
	body, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	result, e := SendToSocket(address, "change-callbacks", body)
	if e != nil {
		return ""
	}
	return string(result)
}

func sendState(address string, entryPoint string, src string) string {
	req := models.StateRequest{
		Type:       "change-state",
		EntryPoint: entryPoint,
		Src:        src,
	}
	body, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	result, e := SendToSocket(address, "change-state", body)
	if e != nil {
		return ""
	}
	return string(result)
}

func sendInfo(address string) string {
	req := models.InfoRequest{
		Type: "change-info",
	}
	body, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	result, e := SendToSocket(address, "hooks-info", body)
	if e != nil {
		return ""
	}
	return string(result)
}
