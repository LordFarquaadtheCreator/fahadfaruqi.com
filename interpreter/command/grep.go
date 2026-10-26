package command

import (
	"fmt"
	"regexp"
	"strings"

	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
)

// grep searches a file's lines. -i ignores case, -n numbers the matches, -v inverts, -c counts
// instead of printing. No match prints nothing at all, which is what grep does when it finds none.
func Grep(args []string, w world.World, s *session.Session) response.Response {
	flags, operands := splitFlags(args)
	if len(operands) < 2 {
		return response.Error("grep", "usage: grep [-i] [-n] [-v] [-c] <pattern> <file...>", s.CWD)
	}

	pattern := operands[0]
	if hasFlag(flags, "-i") {
		pattern = "(?i)" + pattern
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return response.Error("grep", "invalid pattern: "+operands[0], s.CWD)
	}

	numbered := hasFlag(flags, "-n")
	invert := hasFlag(flags, "-v")
	countOnly := hasFlag(flags, "-c")
	files := operands[1:]
	showName := len(files) > 1

	var out []string
	total := 0
	for _, name := range files {
		node, bad, ok := file(name, w, s, "grep")
		if !ok {
			return bad
		}
		count := 0
		for i, line := range contentLines(node.Content) {
			if re.MatchString(line) == invert {
				continue
			}
			count++
			if countOnly {
				continue
			}
			entry := line
			if numbered {
				entry = fmt.Sprintf("%d:%s", i+1, entry)
			}
			if showName {
				entry = name + ":" + entry
			}
			out = append(out, entry)
		}
		total += count
		if countOnly {
			out = append(out, fmt.Sprintf("%s:%d", name, count))
		}
	}

	if countOnly {
		if len(files) == 1 {
			return textResult(fmt.Sprint(total), s)
		}
		return textResult(strings.Join(out, "\n"), s)
	}
	if len(out) == 0 {
		return response.Response{Type: response.Noop, CWD: s.CWD}
	}
	return textResult(strings.Join(out, "\n"), s)
}
