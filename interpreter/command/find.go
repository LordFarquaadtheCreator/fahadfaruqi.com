package command

import (
	"path"
	"sort"
	"strings"

	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
)

// Find walks a directory and prints what matches. -name takes a shell-style pattern and is matched
// against the last part of each path, as find does.
func Find(args []string, w world.World, s *session.Session) response.Response {
	pattern, rest := flagValue(args, "-name", "--name")
	_, operands := splitFlags(rest)

	start := s.CWD
	if len(operands) > 0 {
		resolved, err := w.FS.Resolve(s.CWD, operands[0], w.Env)
		if err != nil {
			return response.Error("find", "no such file or directory: "+operands[0], s.CWD)
		}
		start = resolved
	}

	var found []string
	walk(w, start, func(p string) {
		if pattern == "" {
			found = append(found, p)
			return
		}
		ok, err := path.Match(pattern, path.Base(p))
		if err != nil {
			return
		}
		if ok {
			found = append(found, p)
		}
	})

	if len(found) == 0 {
		return response.Response{Type: response.Noop, CWD: s.CWD}
	}
	sort.Strings(found)
	return textResult(strings.Join(found, "\n"), s)
}

// walk visits every path at or under start, deepest last, in the order the filesystem lists them.
func walk(w world.World, start string, visit func(string)) {
	visit(start)
	node, err := w.FS.Stat(start)
	if err != nil {
		return
	}
	for _, child := range node.Children {
		walk(w, child, visit)
	}
}
