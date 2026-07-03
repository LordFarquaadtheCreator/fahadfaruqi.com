package fs

import "errors"

var ErrNotExist = errors.New("path does not exist")
var ErrNotDir = errors.New("not a directory")

type FileSystem struct {
	nodes map[string]Node
}

// Creates and returns a new FileSystem with the given nodes
func NewFileSystem(nodes map[string]Node) FileSystem {
	return FileSystem{nodes: nodes}
}

// Returns the Node at the given path or an error if it doesn't exist
func (f FileSystem) Stat(absPath string) (Node, error) {
	node, ok := f.nodes[absPath]
	if !ok {
		return Node{}, ErrNotExist
	}
	return node, nil
}

// Returns the children of the directory at absPath, optionally excluding hidden files
func (f FileSystem) Children(absPath string, showHidden bool) ([]Node, []string, error) {
	parent, err := f.Stat(absPath)
	if err != nil {
		return nil, nil, err
	}
	if parent.Type != Dir {
		return nil, nil, ErrNotDir
	}

	var result []Node
	var paths []string
	for _, childPath := range parent.Children {
		child, err := f.Stat(childPath)
		if err != nil {
			continue
		}
		if !showHidden && child.Hidden {
			continue
		}
		result = append(result, child)
		paths = append(paths, childPath)
	}
	return result, paths, nil
}

// Returns true if the path exists in the file system
func (f FileSystem) Exists(absPath string) bool {
	_, ok := f.nodes[absPath]
	return ok
}

// Returns true if the path exists and is a directory
func (f FileSystem) IsDir(absPath string) bool {
	node, ok := f.nodes[absPath]
	return ok && node.Type == Dir
}

// Returns true if the path exists and is a file
func (f FileSystem) IsFile(absPath string) bool {
	node, ok := f.nodes[absPath]
	return ok && node.Type == File
}
