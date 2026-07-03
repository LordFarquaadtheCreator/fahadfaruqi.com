package command

import (
	"strings"
	"interpreter/response"
	"interpreter/world"
	"interpreter/session"
)

func Echo(args []string, w world.World, s *session.Session) response.Response {
	text := strings.Join(args, " ")
	return response.Response{
		Type:    response.Echo,
		Payload: response.TextPayload{Text: text},
		CWD:     s.CWD,
	}
}
