package command

import (
	"fmt"
	"strconv"
	"strings"

	"interpreter/fs"
	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
)

// textResult is the shell's plain-text answer: the terminal renders it as preformatted output.
func textResult(text string, s *session.Session) response.Response {
	return response.Response{
		Type:    response.Echo,
		Payload: response.TextPayload{Text: text},
		CWD:     s.CWD,
	}
}

// splitFlags separates options from operands the way the handlers here read them: anything that
// looks like a switch is a flag, everything else names a file or a pattern.
func splitFlags(args []string) (flags, operands []string) {
	for _, a := range args {
		if len(a) > 1 && a[0] == '-' && !isNegativeNumber(a) {
			flags = append(flags, a)
		} else {
			operands = append(operands, a)
		}
	}
	return flags, operands
}

func isNegativeNumber(a string) bool {
	if len(a) < 2 || a[0] != '-' {
		return false
	}
	_, err := strconv.Atoi(a[1:])
	return err == nil
}

func hasFlag(flags []string, names ...string) bool {
	for _, f := range flags {
		for _, name := range names {
			if f == name {
				return true
			}
		}
	}
	return false
}

// intFlag reads a counted option — -n 5, -n5, --lines=5 — out of the arguments and leaves the rest,
// so the caller can then tell flags from operands. A missing value is an error rather than a silent
// default: head -n with nothing after it is a typo, not a request for ten lines.
func intFlag(args []string, def int, short, long string) (int, []string, error) {
	value, rest := flagValue(args, short, long)
	if value == "" {
		return def, rest, nil
	}
	n, err := strconv.Atoi(value)
	if err != nil || n < 0 {
		return 0, rest, fmt.Errorf("invalid number: %s", value)
	}
	return n, rest, nil
}

// flagValue finds an option that takes a value and returns that value with the option removed.
func flagValue(args []string, short, long string) (string, []string) {
	for i, a := range args {
		var value string
		matched := false
		switch {
		case a == short || a == long:
			if i+1 < len(args) {
				value = args[i+1]
				return value, append(append([]string{}, args[:i]...), args[i+2:]...)
			}
			return "", append(append([]string{}, args[:i]...), args[i+1:]...)
		case long != "" && strings.HasPrefix(a, long+"="):
			value = a[len(long)+1:]
			matched = true
		case len(a) > len(short) && strings.HasPrefix(a, short) && !strings.HasPrefix(a, "--"):
			value = a[len(short):]
			matched = true
		}
		if matched {
			return value, append(append([]string{}, args[:i]...), args[i+1:]...)
		}
	}
	return "", args
}

// file resolves the one file an argument names, with the error a command should report when it is
// missing and a complaint when it is a folder.
func file(name string, w world.World, s *session.Session, cmd string) (fs.Node, response.Response, bool) {
	abs, err := w.FS.Resolve(s.CWD, name, w.Env)
	if err != nil {
		return fs.Node{}, response.Error(cmd, "no such file or directory: "+name, s.CWD), false
	}
	node, err := w.FS.Stat(abs)
	if err != nil {
		return fs.Node{}, response.Error(cmd, "no such file or directory: "+name, s.CWD), false
	}
	if node.Type.IsDir() {
		return fs.Node{}, response.Error(cmd, "is a directory: "+name, s.CWD), false
	}
	return node, response.Response{}, true
}

// contentLines splits a file into lines, dropping the empty last element a trailing newline leaves.
func contentLines(content string) []string {
	if content == "" {
		return nil
	}
	return strings.Split(strings.TrimSuffix(content, "\n"), "\n")
}
