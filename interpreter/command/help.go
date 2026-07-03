package command

import (
	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
	"strings"
)

func Help(args []string, w world.World, s *session.Session) response.Response {
	text := `Available commands:
  ls [path]        list directory contents
  cd <path>        change directory
  cat <file>       display file contents
  open <path>      open file or directory in app
  pwd              print working directory
  echo <text>      display text
  env              display environment variables
  history          display command history
  apps             list available applications
  clear            clear terminal
  help             display this help message`
	return response.Response{
		Type:    response.Echo,
		Payload: response.TextPayload{Text: strings.TrimSpace(text)},
		CWD:     s.CWD,
	}
}
