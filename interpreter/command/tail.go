package command

import (
	"strings"

	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
)

func Tail(args []string, w world.World, s *session.Session) response.Response {
	count, rest, err := intFlag(args, 10, "-n", "--lines")
	if err != nil {
		return response.Error("tail", err.Error(), s.CWD)
	}
	_, operands := splitFlags(rest)
	if len(operands) == 0 {
		return response.Error("tail", "usage: tail [-n lines] <file>", s.CWD)
	}

	node, bad, ok := file(operands[0], w, s, "tail")
	if !ok {
		return bad
	}

	lines := contentLines(node.Content)
	if count < len(lines) {
		lines = lines[len(lines)-count:]
	}
	return textResult(strings.Join(lines, "\n"), s)
}
