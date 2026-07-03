package command

import (
	"interpreter/apps"
	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
	"path"
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

	// If directory, open in Finder
	if w.FS.IsDir(absPath) {
		w.Finder.Spawn(absPath)
		return response.Response{
			Type:    response.Open,
			Payload: response.TextPayload{Text: "Opened " + absPath},
			CWD:     s.CWD,
		}
	}

	// File — resolve app association and spawn viewer
	appID := apps.AppForFile(path.Base(absPath))
	if appID == "" {
		return response.Error("open", "no app associated with file: "+absPath, s.CWD)
	}

	if appID == "viewer-image" || appID == "viewer-text" {
		w.Viewer.SpawnForFile(absPath)
	} else {
		// Placeholder or editor — spawn with that app
		app, ok := apps.Get(appID)
		if !ok {
			return response.Error("open", "unknown app: "+appID, s.CWD)
		}
		w.WM.Spawn(string(app.Type), app.ID, path.Base(absPath), app.Width, app.Height)
	}

	return response.Response{
		Type:    response.Open,
		Payload: response.TextPayload{Text: "Opened " + absPath},
		CWD:     s.CWD,
	}
}
