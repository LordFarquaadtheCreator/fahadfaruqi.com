package command

import (
	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
	"strings"
)

func Echo(args []string, w world.World, s *session.Session) response.Response {
	text := strings.Join(args, " ")
	return response.Response{
		Type:    response.Echo,
		Payload: response.TextPayload{Text: text},
		CWD:     s.CWD,
	}
}
