package shell

import (
	"strings"

	"interpreter/command"
	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
)

func Run(input string, w world.World, s *session.Session) string {
	input = strings.TrimSpace(input)

	if input == "" {
		return response.Marshal(response.Response{Type: response.Clear, CWD: s.CWD})
	}

	s.AppendHistory(input)
	return response.Marshal(eval(input, w, s, 0))
}

// eval runs one line. depth counts how many scripts deep the line is, so that a script run from a
// script can be stopped before it runs forever.
func eval(input string, w world.World, s *session.Session, depth int) response.Response {
	input = strings.TrimSpace(input)

	if input == "" || strings.HasPrefix(input, "#") {
		return response.Response{Type: response.Noop, CWD: s.CWD}
	}

	if k, v, ok := IsVarAssignment(input); ok {
		s.SetVar(k, v)
		return response.Response{Type: response.VarSet, CWD: s.CWD}
	}

	tokens := Tokenize(input)
	if len(tokens) == 0 {
		return response.Response{Type: response.Clear, CWD: s.CWD}
	}

	name := tokens[0]
	args := Expand(tokens[1:], w, s)

	switch name {
	case "source", ".":
		return source(args, w, s, depth)
	}

	// A slash means the visitor named a file, not a command: ./gallery.sh, ../bin/x.sh,
	// ~/Desktop/gallery.sh. A shell only runs something from a path when it is written this way.
	if strings.Contains(name, "/") {
		return executable(name, w, s, depth)
	}

	fn, ok := command.Registry[name]
	if !ok {
		return response.Error(name, "command not found: "+name, s.CWD)
	}

	return fn(args, w, s)
}
