package command

import (
	"fmt"
	"maps"
	"path/filepath"
	"sort"

	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
)

func LS(args []string, w world.World, s *session.Session) response.Response {
	paths := make([]string, 0)
	flags := make([]string, 0)

	// Phase 1: Separate path arguments from flags
	for _, arg := range args {
		if arg == "" || arg[0] == '-' {
			flags = append(flags, arg)
		} else {
			paths = append(paths, arg)
		}
	}

	// Resolve paths to their full absolute form
	targetPath := s.CWD
	for _, path := range paths {
		resolved, err := w.FS.Resolve(s.CWD, path, w.Env)
		if err != nil {
			return response.Error("ls", err.Error(), s.CWD)
		}
		targetPath = resolved
	}

	options := &response.LsOptions{}

	// Phase 2: Parse flags (switch handles all supported flags)
	for _, arg := range flags {
		switch arg {
		case "-a", "--all":
			options.ShowHidden = true

		case "-A", "--no-group-name":
			options.ShowHidden = true

		case "-l", "--long":
			options.LongFormat = true

		case "-h", "--human-readable", "-H":
			options.HumanSize = true

		case "-R", "--recursive", "--recurse":
			options.Recursing = true

		case "-t", "--time=change", "--sort=time":
			options.SortByTime = true

		case "-u", "--time=access":
			options.SortByTime = true

		case "-c", "--time=access":
			options.SortByTime = true

		case "-S", "--size=size", "--sort=size":
			options.SortBySize = true

		case "-r", "--reverse":
			options.SortByReverse = true

		case "-f", "--full-time":
			options.FullTime = true

		case "-1", "--oneline":
			options.Oneline = true

		case "-d", "--directory":
			options.DirectoryOnly = true

		case "-G", "--color=auto", "--color=always", "--color=never":
			options.ColorMode = arg

		default:
			return response.Error("ls", fmt.Sprintf("unknown flag %s", arg), s.CWD)
		}
	}

	// Get children
	children, pathsList, err := w.FS.Children(targetPath, options.ShowHidden)
	if err != nil {
		return response.Error("ls", err.Error(), s.CWD)
	}

	// Sort if needed
	if options.SortBySize && len(children) > 1 {
		sort.Slice(pathsList, func(i, j int) bool {
			return children[i].Meta.Size > children[j].Meta.Size
		})
	} else if options.SortByTime && len(children) > 1 {
		sort.Slice(pathsList, func(i, j int) bool {
			return children[i].Meta.Modified < children[j].Meta.Modified // newest first
		})
	}

	// Build entries
	entries := make([]response.FSEntry, len(children))
	for i, child := range children {
		entries[i] = response.FSEntry{
			Name:     filepath.Base(pathsList[i]),
			Type:     string(child.Type),
			Size:     child.Meta.Size,
			Modified: child.Meta.Modified,
			Hidden:   child.Hidden,
		}

		if options.LongFormat {
			entries[i].Permission = "-rwxr-xr-x"
			entries[i].Owner = "fahad"
			entries[i].Group = "fahadfaruqi.com"
		}
	}

	merged := make(map[string]string)
	maps.Copy(merged, w.Env)
	maps.Copy(merged, s.Vars)

	return response.Response{
		Type:    response.LS,
		Payload: response.LSPayload{Entries: entries, LongFormat: options.LongFormat, HumanSize: options.HumanSize},
		CWD:     s.CWD,
	}
}
