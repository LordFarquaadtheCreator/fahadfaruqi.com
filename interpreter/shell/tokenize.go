package shell

import (
	"strings"
	"unicode"
)

func Tokenize(input string) []string {
	var tokens []string
	var current strings.Builder
	inQuote := false
	quoteChar := rune(0)

	for i, r := range input {
		switch {
		case r == '"' || r == '\'':
			if !inQuote {
				inQuote = true
				quoteChar = r
			} else if r == quoteChar {
				inQuote = false
				quoteChar = 0
			} else {
				current.WriteRune(r)
			}
		case unicode.IsSpace(r):
			if inQuote {
				current.WriteRune(r)
			} else {
				if current.Len() > 0 {
					tokens = append(tokens, current.String())
					current.Reset()
				}
			}
		default:
			current.WriteRune(r)
		}

		// Handle comment at end
		if r == '#' && !inQuote && i > 0 && unicode.IsSpace(rune(input[i-1])) {
			break
		}
	}

	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}

	return tokens
}
