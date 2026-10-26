package command

import (
	"strings"

	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
)

func Head(args []string, w world.World, s *session.Session) response.Response {
	count, rest, err := intFlag(args, 10, "-n", "--lines")
	if err != nil {
		return response.Error("head", err.Error(), s.CWD)
	}
	_, operands := splitFlags(rest)
	if len(operands) == 0 {
		return response.Error("head", "usage: head [-n lines] <file>", s.CWD)
	}

	node, bad, ok := file(operands[0], w, s, "head")
	if !ok {
		return bad
	}

	lines := contentLines(node.Content)
	if count < len(lines) {
		lines = lines[:count]
	}
	return textResult(strings.Join(lines, "\n"), s)
}
