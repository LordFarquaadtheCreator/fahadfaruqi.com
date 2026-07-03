package command

import (
	"interpreter/response"
	"interpreter/world"
	"interpreter/session"
)

func PWD(args []string, w world.World, s *session.Session) response.Response {
	return response.Response{
		Type:    response.PWD,
		Payload: response.TextPayload{Text: s.CWD},
		CWD:     s.CWD,
	}
}
