package command

import (
	"interpreter/fs"
	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
)

func CD(args []string, w world.World, s *session.Session) response.Response {
	if len(args) == 0 {
		s.SetCWD(w.Env["HOME"])
		return response.Response{Type: response.Noop, Payload: nil, CWD: s.CWD}
	}
	if len(args) > 1 {
		return response.Error("cd", "too many arguments", s.CWD)
	}

	target := args[0]
	resolved, err := w.FS.Resolve(s.CWD, target, w.Env)
	if err != nil {
		return response.Error("cd", err.Error(), s.CWD)
	}

	node, err := w.FS.Stat(resolved)
	if err != nil {
		return response.Error("cd", err.Error(), s.CWD)
	}

	if node.Type != fs.Dir {
		return response.Error("cd", "not a directory: "+target, s.CWD)
	}

	s.SetCWD(resolved)
	return response.Response{Type: response.Noop, Payload: nil, CWD: s.CWD}
}
