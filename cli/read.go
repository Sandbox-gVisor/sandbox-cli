package cli

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"sandbox-cli/models"
)

func ReadFile(filename string) []byte {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	return b
}

func extractSyscallFunctions(jsSource string) []models.CallbackDto {
	callbacks := make([]models.CallbackDto, 0)

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

		callback := models.CallbackDto{
			CallbackSource: jsSource,
			EntryPoint:     functionName,
			Sysno:          sysno,
			Type:           callbackType,
		}
		callbacks = append(callbacks, callback)
	}

	return callbacks
}

func ReadDtos(content []byte) []models.CallbackDto {
	return extractSyscallFunctions(string(content))
}

func ReadSource(content []byte) string {
	// TODO tut budet chto-to pointeresnee
	return string(content)
}
