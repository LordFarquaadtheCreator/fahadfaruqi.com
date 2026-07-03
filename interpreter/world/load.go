package world

import (
	_ "embed"
	"encoding/json"
	"interpreter/finder"
	"interpreter/fs"
	"interpreter/viewer"
	"interpreter/wm"
	"strings"
)

//go:embed data.json
var raw []byte

func Load() World {
	var data map[string]jsonNode
	if err := json.Unmarshal(raw, &data); err != nil {
		panic("failed to parse data.json: " + err.Error())
	}

	nodes := make(map[string]fs.Node)
	env := map[string]string{
		"HOME":  "~",
		"USER":  "alex",
		"SHELL": "/bin/zsh",
	}

	// First pass: create all nodes
	for path, jn := range data {
		node := fs.Node{
			Type:     fs.NodeType(jn.Type),
			Children: jn.Children,
			Content:  jn.Content,
			Meta: fs.Meta{
				Size:      int64(len(jn.Content)),
				Modified:  "2024-01-01",
				ImagePath: jn.ImagePath,
				MimeType:  jn.MimeType,
			},
			Hidden: jn.Hidden,
		}
		nodes[path] = node
	}

	// Validate: every child must exist
	for path, node := range nodes {
		if node.Type == fs.Dir {
			for _, childPath := range node.Children {
				if _, ok := nodes[childPath]; !ok {
					panic("orphan child: " + childPath + " referenced from " + path)
				}
			}
		}
	}

	// Validate: no orphan nodes (every non-root node must be referenced as a child)
	for path := range nodes {
		if path == "~" {
			continue
		}
		found := false
		for _, node := range nodes {
			if node.Type == fs.Dir {
				for _, childPath := range node.Children {
					if childPath == path {
						found = true
						break
					}
				}
			}
			if found {
				break
			}
		}
		if !found {
			panic("orphan node: " + path + " not referenced as a child")
		}
	}

	// Validate: hidden flag derived from name
	for path, node := range nodes {
		expectedHidden := strings.HasPrefix(path, "~/.")
		if node.Hidden != expectedHidden {
			panic("hidden flag mismatch for " + path)
		}
	}

	filesystem := fs.NewFileSystem(nodes)
	manager := wm.NewManager()

	w := World{
		FS:     filesystem,
		Env:    env,
		WM:     manager,
		Finder: finder.NewFinderManager(manager, filesystem),
		Viewer: viewer.NewViewerManager(manager, filesystem),
	}
	return w
}

type jsonNode struct {
	Type      string   `json:"type"`
	Children  []string `json:"children"`
	Content   string   `json:"content"`
	Hidden    bool     `json:"hidden"`
	ImagePath string   `json:"imagePath"`
	MimeType  string   `json:"mimeType"`
}
