package prettyoutput

import (
	"regexp"
	"strings"
)

func HighlightJsSyntax(jsSource string) string {
	type colorScheme struct {
		patterns []string
		color    int
	}

	schemes := []colorScheme{
		{
			patterns: []string{"function", "if", "return", "const", "let", "var"},
			color:    GreenColorText,
		},
		{
			patterns: []string{`\(`, `\)`},
			color:    OrangeColorText,
		},
		{
			patterns: []string{`"`, "{", "}"},
			color:    RedColorText,
		},
	}

	removeEscaping := func(str string) string {
		return strings.Replace(str, `\`, "", -1)
	}

	for _, scheme := range schemes {
		for _, pattern := range scheme.patterns {
			reg := regexp.MustCompile(pattern)
			jsSource = reg.ReplaceAllString(jsSource, MakeTextBoldAndColored(removeEscaping(pattern), scheme.color))
		}
	}

	return jsSource
}
