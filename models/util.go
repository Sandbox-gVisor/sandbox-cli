package models

import (
	"fmt"
	"os"
)

type CliError struct {
	Message string
	Cause   error
}

func (c CliError) Error() string {
	return fmt.Sprintf("%s, cause: %v", c.Message, c.Cause)
}

func MakeCliError(message string, cause error) CliError {
	return CliError{Message: message, Cause: cause}
}

func ReadFile(filename string) []byte {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	return b
}
