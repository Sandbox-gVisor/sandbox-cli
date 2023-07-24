package models

import "fmt"

func makeTextBold(str string) string {
	return fmt.Sprintf("\033[1m%s\033[0m", str)
}

func makeTextUnderline(str string) string {
	return fmt.Sprintf("\033[4m%s\033[0m", str)
}

func makeTextHighlight(str string) string {
	return fmt.Sprintf("\u001B[7mS\b%s\u001B[0m", str)
}
