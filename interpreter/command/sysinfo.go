package command

import (
	"strings"

	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
)

// Whoami names the visitor the shell is running as, from the environment the world was built with.
func Whoami(args []string, w world.World, s *session.Session) response.Response {
	return textResult(envValue("USER", w, s), s)
}

// Hostname names the machine, which is also the second half of the shell's prompt.
func Hostname(args []string, w world.World, s *session.Session) response.Response {
	return textResult(envValue("HOSTNAME", w, s), s)
}

// Uname reports the system; -a adds the host and the kernel details.
func Uname(args []string, w world.World, s *session.Session) response.Response {
	flags, _ := splitFlags(args)
	if !hasFlag(flags, "-a", "--all") {
		return textResult("Darwin", s)
	}
	host := envValue("HOSTNAME", w, s)
	return textResult(strings.Join([]string{"Darwin", host, "24.0.0", "arm64", "Darwin Kernel Version 24.0.0"}, " "), s)
}

// envValue reads a variable the way the shell does: what the session set first, then the world's.
func envValue(name string, w world.World, s *session.Session) string {
	if v, ok := s.Vars[name]; ok {
		return v
	}
	if v, ok := w.Env[name]; ok {
		return v
	}
	return "unknown"
}
