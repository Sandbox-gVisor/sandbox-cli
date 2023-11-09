package commands

import (
	"os"
)

func ReadFile(filename string) []byte {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	return b
}
