package fs

import (
	"path/filepath"
)

// Resolves all inputs against cwd, raises error if path is invalid
// Assumes cwd is clean
func (f FileSystem) Resolve(cwd, input string, env map[string]string) (string, error) {
	// Clean user input
	cleanedInput := filepath.Clean(input)

	// Join and validate abs path
	joinedAbsPath := filepath.Join(cwd, cleanedInput)

	if !f.Exists(joinedAbsPath) {
		return "", ErrNotExist
	}

	return joinedAbsPath, nil
}
