package command

import (
	"maps"

	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
)

func Env(args []string, w world.World, s *session.Session) response.Response {
	merged := make(map[string]string)
	maps.Copy(merged, w.Env)
	maps.Copy(merged, s.Vars)

	return response.Response{
		Type:    response.Env,
		Payload: response.EnvPayload{Vars: merged},
		CWD:     s.CWD,
	}
}
