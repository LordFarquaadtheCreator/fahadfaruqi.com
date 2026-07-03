package command

import (
	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
)

func Open(args []string, w world.World, s *session.Session) response.Response {
	if len(args) == 0 {
		return response.Error("open", "usage: open <path>", s.CWD)
	}
	target := args[0]
	absPath, err := w.FS.Resolve(s.CWD, target, w.Env)
	if err != nil {
		return response.Error("open", err.Error(), s.CWD)
	}
	if !w.FS.Exists(absPath) {
		return response.Error("open", "path does not exist: "+absPath, s.CWD)
	}
	if _, err := w.Finder.Spawn(absPath, w.FS); err != nil {
		return response.Error("open", err.Error(), s.CWD)
	}
	return response.Response{
		Type:    response.Open,
		Payload: response.TextPayload{Text: "Opened " + absPath},
		CWD:     s.CWD,
	}
}
