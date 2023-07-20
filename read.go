package main

import (
	"encoding/json"
	"os"
)

func ReadFile(filename string) []byte {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	return b
}

func ReadDtos(content []byte) []CallbackDto {
	var callbacks []CallbackDto
	json.Unmarshal(content, &callbacks)
	return callbacks
}
