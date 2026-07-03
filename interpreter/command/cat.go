package command

import (
	"interpreter/fs"
	"interpreter/response"
	"interpreter/world"
	"interpreter/session"
)

/*
 * cat should only have one argument
 */

func Cat(args []string, w world.World, s *session.Session) response.Response {
	if len(args) == 0 {
		return response.Error("cat", "missing file argument", s.CWD)
	}
	if len(args) > 1 {
		return response.Error("cat", "too many arguments",
		s.CWD)
	}

	// resolve path
	target := args[0]
	resolvedPath, err := w.FS.Resolve(s.CWD, target, w.Env)
	if err != nil {
		return response.Error("cat", err.Error(), s.CWD)
	}

	node, err := w.FS.Stat(resolvedPath)
	if err != nil {
		return response.Error("cat", err.Error(), s.CWD)
	}

	if node.Type != fs.File {
		return response.Error("cat", "not a file: "+target, s.CWD)
	}

	return response.Response{
		Type:    response.Cat,
		Payload: response.CatPayload{Name: target, Content: node.Content},
		CWD:     s.CWD,
	}
}
