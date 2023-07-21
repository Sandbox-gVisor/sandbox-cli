package cli

import (
	"encoding/json"
	"os"

	"sandbox-cli/models"
)

func ReadFile(filename string) []byte {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	return b
}

func ReadDtos(content []byte) []models.CallbackDto {
	var callbacks []models.CallbackDto
	if err := json.Unmarshal(content, &callbacks); err != nil {
		return []models.CallbackDto{}
	}
	return callbacks
}

func ReadSource(content []byte) string {
	// TODO tut budet chto-to pointeresnee
	return string(content)
}
