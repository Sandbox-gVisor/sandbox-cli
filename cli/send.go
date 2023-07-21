package cli

import (
	"encoding/json"
	"fmt"
	"sandbox-cli/models"
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
	result, e := SendToSocket(address, body)
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
	result, e := SendToSocket(address, body)
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
	result, e := SendToSocket(address, body)
	if e != nil {
		return ""
	}
	return string(result)
}

func sendGet(address string) string {
	req := models.GetRequest{
		Type: "current-callbacks",
	}
	body, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	result, e := SendToSocket(address, body)
	if e != nil {
		return ""
	}
	return string(result)
}

func sendDelete(address string, options string, sysno int, callbackType string) string {
	var list []models.DeleteCallbackJson
	if options != "all" {
		list = append(list, models.DeleteCallbackJson{
			Type:  callbackType,
			Sysno: sysno,
		})
	}
	req := models.DeleteRequest{
		Type:    "unregister-callbacks",
		Options: options,
		List:    list,
	}
	body, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	result, e := SendToSocket(address, body)
	if e != nil {
		return ""
	}
	return string(result)
}
