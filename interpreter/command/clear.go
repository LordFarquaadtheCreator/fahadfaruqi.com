package command

import (
	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
)

func Clear(args []string, w world.World, s *session.Session) response.Response {
	return response.Response{
		Type:    response.Clear,
		Payload: nil,
		CWD:     s.CWD,
	}
}
