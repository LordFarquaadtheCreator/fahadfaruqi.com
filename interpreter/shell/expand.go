package shell

import (
	"strings"
	"interpreter/world"
	"interpreter/session"
)

func Expand(tokens []string, w world.World, s *session.Session) []string {
	var expanded []string
	for _, token := range tokens {
		if strings.HasPrefix(token, "$") {
			varName := token[1:]
			// Check session vars first
			if val, ok := s.Vars[varName]; ok {
				expanded = append(expanded, val)
				continue
			}
			// Fall back to world env
			if val, ok := w.Env[varName]; ok {
				expanded = append(expanded, val)
				continue
			}
			// Keep original if not found
			expanded = append(expanded, token)
		} else {
			expanded = append(expanded, token)
		}
	}
	return expanded
}

func IsVarAssignment(input string) (string, string, bool) {
	parts := strings.SplitN(input, "=", 2)
	if len(parts) != 2 {
		return "", "", false
	}
	key := strings.TrimSpace(parts[0])
	val := strings.TrimSpace(parts[1])
	return key, val, true
}
