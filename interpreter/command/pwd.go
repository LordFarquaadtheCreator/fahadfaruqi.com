package command

import (
	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
)

func PWD(args []string, w world.World, s *session.Session) response.Response {
	return response.Response{
		Type:    response.PWD,
		Payload: response.TextPayload{Text: s.CWD},
		CWD:     s.CWD,
	}
}
