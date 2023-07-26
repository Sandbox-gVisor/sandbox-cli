package models

import (
	"fmt"
	"regexp"
	"strconv"
)

type CallbackDto struct {
	Sysno          int    `json:"sysno"`
	EntryPoint     string `json:"entry-point"`
	CallbackSource string `json:"source"`
	Type           string `json:"type"`
}

type ChangeRequest struct {
	Callbacks []CallbackDto `json:"callbacks"`
}

func extractSyscallFunctions(jsSource string) []CallbackDto {
	callbacks := make([]CallbackDto, 0)

	re := regexp.MustCompile(`function\s+(syscall_(\w+)_(\d+))`)
	matches := re.FindAllStringSubmatch(jsSource, -1)

	for _, match := range matches {
		functionName := match[1]
		callbackType := match[2]
		sysno, err := strconv.Atoi(match[3])
		if err != nil {
			fmt.Printf("Error parsing sysno for function %s: %v\n", functionName, err)
			continue
		}

		callback := CallbackDto{
			CallbackSource: jsSource,
			EntryPoint:     functionName,
			Sysno:          sysno,
			Type:           callbackType,
		}
		callbacks = append(callbacks, callback)
	}

	return callbacks
}

func MakeChangeCallbacksRequest(jsSourceFileName string) *Request {
	dtos := extractSyscallFunctions(string(ReadFile(jsSourceFileName)))
	
	req := &Request{
		Type:    "change-callbacks",
		Payload: ChangeRequest{Callbacks: dtos},
	}
	return req
}
