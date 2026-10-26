package command

import (
	"strings"

	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
)

func Help(args []string, w world.World, s *session.Session) response.Response {
	text := `Available commands:
  ls [path]              list directory contents
  cd <path>              change directory
  pwd                    print working directory
  find [path] [-name p]  find files by name
  tree [path]            show the folders as a tree

  cat <file>             display file contents
  head [-n N] <file>     first lines of a file
  tail [-n N] <file>     last lines of a file
  wc [-l|-w|-c] <file>   count lines, words and bytes
  grep [-i] [-n] [-v] [-c] <pattern> <file>
                         search a file's lines
  stat <path>            what the filesystem knows about a path

  open <path|url>        open a file, a folder, or a page in the browser
  apps                   list applications
  which <command>        where a command comes from
  env                    environment variables
  date                   the time
  whoami                 who the shell is running as
  hostname               this machine's name
  uname [-a]             the system

  echo <text>            display text
  history                command history
  clear                  clear the terminal
  help                   this message

  ./<script.sh>          run a shell script
  source <file>          run a script in this shell, keeping its variables`
	return response.Response{
		Type:    response.Echo,
		Payload: response.TextPayload{Text: strings.TrimSpace(text)},
		CWD:     s.CWD,
	}
}
