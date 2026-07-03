package command

import (
	"interpreter/apps"
	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
	"strings"
)

func Apps(args []string, w world.World, s *session.Session) response.Response {
	var names []string
	for _, id := range apps.DockOrder {
		if a, ok := apps.Get(id); ok {
			names = append(names, a.Name)
		}
	}
	return response.Response{
		Type:    response.Apps,
		Payload: response.TextPayload{Text: strings.Join(names, "  ")},
		CWD:     s.CWD,
	}
}
