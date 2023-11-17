package errors

import "fmt"

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
