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
  open <path|url>  open a file, a folder, or a page in the browser
  pwd              print working directory
  echo <text>      display text
  env              display environment variables
  history          display command history
  apps             list available applications
  clear            clear terminal
  help             display this help message
  ./<script.sh>    run a shell script
  source <file>    run a shell script in this terminal, keeping its variables`
	return response.Response{
		Type:    response.Echo,
		Payload: response.TextPayload{Text: strings.TrimSpace(text)},
		CWD:     s.CWD,
	}
}
