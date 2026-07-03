package terminal

import (
	"interpreter/fs"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

// Complete returns tab completion candidates for the given input.
func Complete(input, cwd string, filesystem fs.FileSystem) []string {
	if input == "" {
		return nil
	}

	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		return nil
	}

	// Completing a command name
	if len(tokens) == 1 && !strings.Contains(tokens[0], "/") {
		return completeCommand(tokens[0])
	}

	// Completing a path argument
	lastToken := tokens[len(tokens)-1]
	if strings.HasPrefix(lastToken, "-") {
		return nil
	}

	return completePath(lastToken, cwd, filesystem)
}

func completeCommand(partial string) []string {
	commands := []string{"ls", "cd", "cat", "pwd", "echo", "env", "history", "clear", "open", "help", "apps"}
	var matches []string
	for _, cmd := range commands {
		if strings.HasPrefix(cmd, partial) {
			matches = append(matches, cmd)
		}
	}
	sort.Strings(matches)
	return matches
}

func completePath(partial, cwd string, filesystem fs.FileSystem) []string {
	// Split into dir part and prefix
	dirPart := filepath.Dir(partial)
	basePart := filepath.Base(partial)

	// Resolve directory
	var dirPath string
	if strings.HasPrefix(partial, "~/") || partial == "~" {
		dirPath = "~"
		if strings.HasPrefix(partial, "~/") {
			sub := partial[2:]
			if idx := strings.LastIndex(sub, "/"); idx >= 0 {
				dirPath = "~/" + sub[:idx]
				basePart = sub[idx+1:]
			} else {
				dirPath = "~"
				basePart = sub
			}
		}
	} else if strings.HasPrefix(partial, "/") {
		dirPath = "/" 
	} else if dirPart == "." || dirPart == "" {
		dirPath = cwd
	} else if dirPart == ".." {
		dirPath = filepath.Dir(cwd)
	} else {
		dirPath = filepath.Join(cwd, dirPart)
	}

	// Normalize ~ paths
	dirPath = strings.ReplaceAll(dirPath, "~", "~")

	children, paths, err := filesystem.Children(dirPath, true)
	if err != nil {
		return nil
	}

	var matches []string
	for i, child := range children {
		name := path.Base(paths[i])
		if !strings.HasPrefix(strings.ToLower(name), strings.ToLower(basePart)) {
			continue
		}
		// Build completion preserving original path prefix from input
		var completion string
		if strings.HasPrefix(partial, "~/") || partial == "~" {
			completion = dirPart
			if dirPart != "~" && dirPart != "" {
				completion += "/"
			} else if dirPart == "~" {
				completion += "/"
			}
		} else if dirPart == "." || dirPart == "" {
			completion = ""
		} else {
			completion = dirPart + "/"
		}
		completion += name
		if child.Type.IsDir() {
			completion += "/"
		}
		matches = append(matches, completion)
	}

	sort.Strings(matches)
	return matches
}
