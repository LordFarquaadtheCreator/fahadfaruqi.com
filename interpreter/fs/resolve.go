package fs

import (
	"path/filepath"
	"strings"
)

// Resolves all inputs against cwd, raises error if path is invalid
// Assumes cwd is clean
func (f FileSystem) Resolve(cwd, input string, env map[string]string) (string, error) {
	// Expand ~ to HOME
	if input == "~" {
		input = env["HOME"]
	} else if strings.HasPrefix(input, "~/") {
		input = env["HOME"] + input[1:]
	}

	// Clean user input
	cleanedInput := filepath.Clean(input)

	// If already absolute (starts with ~), don't join with cwd
	var joinedAbsPath string
	if strings.HasPrefix(cleanedInput, env["HOME"]) || cleanedInput == env["HOME"] {
		joinedAbsPath = cleanedInput
	} else {
		joinedAbsPath = filepath.Join(cwd, cleanedInput)
	}

	if !f.Exists(joinedAbsPath) {
		return "", ErrNotExist
	}

	return joinedAbsPath, nil
}
