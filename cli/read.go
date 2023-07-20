package cli

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
	if err := json.Unmarshal(content, &callbacks); err != nil {
		return []CallbackDto{}
	}
	return callbacks
}
