package command

import (
	"fmt"
	"strings"

	"interpreter/fs"
	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
)

// Stat reports what the filesystem knows about a path, in the shape `ls -l` uses for the same facts.
func Stat(args []string, w world.World, s *session.Session) response.Response {
	_, operands := splitFlags(args)
	if len(operands) == 0 {
		return response.Error("stat", "usage: stat <path>", s.CWD)
	}

	abs, err := w.FS.Resolve(s.CWD, operands[0], w.Env)
	if err != nil {
		return response.Error("stat", "no such file or directory: "+operands[0], s.CWD)
	}
	node, err := w.FS.Stat(abs)
	if err != nil {
		return response.Error("stat", "no such file or directory: "+operands[0], s.CWD)
	}

	kind := "file"
	if node.Type == fs.Dir {
		kind = "directory"
	}
	lines := []string{
		"path: " + abs,
		"type: " + kind,
		fmt.Sprintf("size: %d", node.Meta.Size),
		"modified: " + node.Meta.Modified,
	}
	if kind == "directory" {
		lines = append(lines, fmt.Sprintf("entries: %d", len(node.Children)))
	}
	return textResult(strings.Join(lines, "\n"), s)
}
