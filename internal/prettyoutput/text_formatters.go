package prettyoutput

import "fmt"

func MakeTextBold(str string) string {
	return fmt.Sprintf("\033[1m%s\033[0m", str)
}

func MakeTextUnderline(str string) string {
	return fmt.Sprintf("\033[4m%s\033[0m", str)
}

func MakeTextHighlight(str string) string {
	return fmt.Sprintf("\u001B[7mS\b%s\u001B[0m", str)
}

func MakeTextColored(str string, color int) string {
	return fmt.Sprintf("\033[%dm%s\033[0m", color, str)
}

const RedColorText = 31
const GreenColorText = 32
const OrangeColorText = 33
const BlueColorText = 34

func MakeTextBoldAndColored(str string, color int) string {
	return MakeTextBold(MakeTextColored(str, color))
}
