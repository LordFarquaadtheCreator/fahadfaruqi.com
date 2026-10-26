package command

import (
	"strings"

	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
)

// treeDepth keeps the listing to something a terminal window can show.
const treeDepth = 4

const treeLines = 120

func Tree(args []string, w world.World, s *session.Session) response.Response {
	_, operands := splitFlags(args)
	start := s.CWD
	if len(operands) > 0 {
		resolved, err := w.FS.Resolve(s.CWD, operands[0], w.Env)
		if err != nil {
			return response.Error("tree", "no such file or directory: "+operands[0], s.CWD)
		}
		start = resolved
	}

	out := []string{start}
	var walkTree func(dir string, prefix string, depth int)
	walkTree = func(dir string, prefix string, depth int) {
		if depth >= treeDepth || len(out) >= treeLines {
			return
		}
		node, err := w.FS.Stat(dir)
		if err != nil {
			return
		}
		for i, child := range node.Children {
			if len(out) >= treeLines {
				return
			}
			last := i == len(node.Children)-1
			branch, indent := "├── ", "│   "
			if last {
				branch, indent = "└── ", "    "
			}
			childNode, err := w.FS.Stat(child)
			if err != nil {
				continue
			}
			name := child[strings.LastIndex(child, "/")+1:]
			if childNode.Type.IsDir() {
				name += "/"
			}
			out = append(out, prefix+branch+name)
			if childNode.Type.IsDir() {
				walkTree(child, prefix+indent, depth+1)
			}
		}
	}
	walkTree(start, "", 0)

	if len(out) >= treeLines {
		out = append(out, "…")
	}
	return textResult(strings.Join(out, "\n"), s)
}
