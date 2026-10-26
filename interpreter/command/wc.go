package command

import (
	"fmt"
	"strings"

	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
)

// WC counts lines, words and bytes, in that order, and prints the file's name after them the way
// the real one does. With any of -l, -w or -c it prints only what was asked for.
func WC(args []string, w world.World, s *session.Session) response.Response {
	flags, operands := splitFlags(args)
	if len(operands) == 0 {
		return response.Error("wc", "usage: wc [-l] [-w] [-c] <file>", s.CWD)
	}

	node, bad, ok := file(operands[0], w, s, "wc")
	if !ok {
		return bad
	}

	lines := len(contentLines(node.Content))
	words := len(strings.Fields(node.Content))
	bytes := len(node.Content)

	selected := hasFlag(flags, "-l") || hasFlag(flags, "-w") || hasFlag(flags, "-c")
	var parts []string
	if !selected || hasFlag(flags, "-l") {
		parts = append(parts, fmt.Sprint(lines))
	}
	if !selected || hasFlag(flags, "-w") {
		parts = append(parts, fmt.Sprint(words))
	}
	if !selected || hasFlag(flags, "-c") {
		parts = append(parts, fmt.Sprint(bytes))
	}
	parts = append(parts, operands[0])

	return textResult(strings.Join(parts, " "), s)
}
