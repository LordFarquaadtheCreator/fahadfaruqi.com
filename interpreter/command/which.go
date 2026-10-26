package command

import (
	"strings"

	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
)

// builtins is what the shell answers to, kept beside the registry rather than derived from it: a
// handler that reads the registry would be an initialisation cycle, so a test holds the two lists
// to each other instead.
var builtins = []string{
	"ls", "cd", "cat", "pwd", "echo", "env", "history", "clear", "open", "apps", "help",
	"head", "tail", "wc", "grep", "find", "tree", "stat", "which", "date", "whoami", "hostname",
	"uname",
}

func isBuiltin(name string) bool {
	for _, b := range builtins {
		if b == name {
			return true
		}
	}
	return false
}

// Which answers where a command comes from. Everything here is a shell builtin, so a known name
// reports the path a builtin would have; an unknown one says so rather than answering nothing.
func Which(args []string, w world.World, s *session.Session) response.Response {
	_, operands := splitFlags(args)
	if len(operands) == 0 {
		return response.Error("which", "usage: which <command>", s.CWD)
	}
	var out []string
	for _, name := range operands {
		if isBuiltin(name) {
			out = append(out, "/usr/bin/"+name)
		} else {
			out = append(out, name+" not found")
		}
	}
	return textResult(strings.Join(out, "\n"), s)
}
