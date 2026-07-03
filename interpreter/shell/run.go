package shell

import (
	"strings"
	"interpreter/command"
	"interpreter/response"
	"interpreter/world"
	"interpreter/session"
)

func Run(input string, w world.World, s *session.Session) string {
	input = strings.TrimSpace(input)

	if input == "" {
		return response.Marshal(response.Response{Type: response.Clear, CWD: s.CWD})
	}

	s.AppendHistory(input)

	if k, v, ok := IsVarAssignment(input); ok {
		s.SetVar(k, v)
		return response.Marshal(response.Response{Type: response.VarSet, CWD: s.CWD})
	}

	tokens := Tokenize(input)
	if len(tokens) == 0 {
		return response.Marshal(response.Response{Type: response.Clear, CWD: s.CWD})
	}

	name, args := tokens[0], tokens[1:]
	args = Expand(args, w, s)

	fn, ok := command.Registry[name]
	if !ok {
		return response.Marshal(response.Error(name, "command not found: "+name, s.CWD))
	}

	return response.Marshal(fn(args, w, s))
}
