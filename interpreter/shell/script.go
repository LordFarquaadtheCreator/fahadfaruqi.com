package shell

import (
	"strings"

	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
)

// maxDepth bounds how far a script may run another script. A file that sources itself would
// otherwise run until the stack gives out; a shell stops too, just at a number it chose.
const maxDepth = 24

// scriptLines returns the lines of a script that the shell should run: comments and blank lines
// are not commands, and a trailing carriage return is an artefact of the file, not an argument.
func scriptLines(content string) []string {
	var lines []string
	for _, raw := range strings.Split(content, "\n") {
		line := strings.TrimSpace(strings.TrimSuffix(raw, "\r"))
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		lines = append(lines, line)
	}
	return lines
}

// executable runs a path the visitor named with a slash in it — ./gallery.sh, ../bin/x.sh. The path
// has to exist, be a file, and be a shell script: those are the only executables this filesystem
// has, so anything else is a permission error, as a shell would report an unexecutable file.
func executable(name string, w world.World, s *session.Session, depth int) response.Response {
	abs, err := w.FS.Resolve(s.CWD, name, w.Env)
	if err != nil {
		return response.Error(name, "no such file or directory: "+name, s.CWD)
	}
	node, err := w.FS.Stat(abs)
	if err != nil {
		return response.Error(name, "no such file or directory: "+name, s.CWD)
	}
	if node.Type.IsDir() {
		return response.Error(name, "is a directory: "+name, s.CWD)
	}
	if !isScript(abs) {
		return response.Error(name, "permission denied: "+name, s.CWD)
	}
	return runScript(abs, node.Content, w, s, depth+1)
}

// source runs a script in the terminal's own shell, as `source file` and `. file` do. It asks for
// read access, not execute — which is the difference between sourcing a file and running it — and
// it takes the file's path as given, so `source gallery.sh` works without a leading ./
func source(args []string, w world.World, s *session.Session, depth int) response.Response {
	if len(args) == 0 {
		return response.Error("source", "usage: source <file>", s.CWD)
	}
	name := args[0]
	abs, err := w.FS.Resolve(s.CWD, name, w.Env)
	if err != nil {
		return response.Error("source", "no such file or directory: "+name, s.CWD)
	}
	node, err := w.FS.Stat(abs)
	if err != nil {
		return response.Error("source", "no such file or directory: "+name, s.CWD)
	}
	if node.Type.IsDir() {
		return response.Error("source", "is a directory: "+name, s.CWD)
	}
	if !isScript(abs) {
		return response.Error("source", "not a shell script: "+name, s.CWD)
	}
	return runScript(abs, node.Content, w, s, depth+1)
}

// runScript evaluates a script's lines in order and reports the last thing that happened, which is
// what the terminal shows beneath the command that started it. A line that fails stops the script:
// these scripts are three lines long and a failure halfway through is worth seeing, not skipping.
func runScript(abs, content string, w world.World, s *session.Session, depth int) response.Response {
	if depth > maxDepth {
		return response.Error("source", "maximum nesting depth reached: "+abs, s.CWD)
	}
	last := response.Response{Type: response.Noop, CWD: s.CWD}
	for _, line := range scriptLines(content) {
		r := eval(line, w, s, depth)
		if r.Type == response.Err {
			return r
		}
		if r.Type != response.Noop {
			last = r
		}
	}
	return last
}

func isScript(path string) bool {
	return strings.HasSuffix(strings.ToLower(path), ".sh")
}
