package command

import (
	"interpreter/response"
	"interpreter/world"
	"interpreter/session"
)

func History(args []string, w world.World, s *session.Session) response.Response {
	return response.Response{
		Type:    response.History,
		Payload: response.HistoryPayload{Entries: s.History},
		CWD:     s.CWD,
	}
}
