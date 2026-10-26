package finder

import (
	"encoding/json"
	"interpreter/assets"
	"interpreter/fs"
	"interpreter/markdown"
	"interpreter/wm"
	"path"
	"strings"
)

// FinderContentVM is the app-specific content for a Finder window.
type FinderContentVM struct {
	Kind          string             `json:"kind"` // "finder"
	ViewMode      string             `json:"viewMode"`
	SelectedEntry string             `json:"selectedEntry"`
	CanBack       bool               `json:"canBack"`
	CanForward    bool               `json:"canForward"`
	Sidebar       []SidebarSectionVM `json:"sidebar"`
	Breadcrumb    []BreadcrumbItemVM `json:"breadcrumb"`
	Content       ContentVM          `json:"content"`
}

type SidebarSectionVM struct {
	Title string          `json:"title"` // "", "Favorites", "Locations"
	Items []SidebarItemVM `json:"items"`
}

type SidebarItemVM struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Icon   string `json:"icon"` // icon path or type
	Active bool   `json:"active"`
}

type ContentVM struct {
	Kind    string    `json:"kind"` // "folder" | "markdown" | "image" | "text"
	Entries []EntryVM `json:"entries,omitempty"`
	HTML    string    `json:"html,omitempty"`
}

type EntryVM struct {
	Name           string `json:"name"`
	Path           string `json:"path"`
	Kind           string `json:"kind"` // "dir" | "md" | "txt" | "png" | "jpeg"
	Icon           string `json:"icon"`
	ImagePath      string `json:"imagePath,omitempty"`
	CompressedPath string `json:"compressedPath,omitempty"`
	Size           int64  `json:"size"`
	Modified       string `json:"modified"`
}

type BreadcrumbItemVM struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Icon string `json:"icon"`
}

// BuildContentVM builds the Finder-specific content for a window.
func BuildContentVM(win *wm.Window, fm *FinderManager, filesystem fs.FileSystem) json.RawMessage {
	s, ok := fm.GetState(win.ID)
	if !ok {
		return json.RawMessage(`{"kind":"finder","content":{"kind":"folder","entries":[]}}`)
	}

	vm := FinderContentVM{
		Kind:          "finder",
		ViewMode:      string(s.ViewMode),
		SelectedEntry: s.SelectedEntry,
		CanBack:       s.CanBack(),
		CanForward:    s.CanForward(),
		Sidebar:       buildSidebar(s.Path, filesystem),
		Breadcrumb:    buildBreadcrumb(s.Path),
	}

	node, err := filesystem.Stat(s.Path)
	if err != nil {
		vm.Content = ContentVM{Kind: "folder", Entries: []EntryVM{}}
	} else if node.Type.IsDir() {
		children, paths, _ := filesystem.Children(s.Path, false)
		var entries []EntryVM
		for i, child := range children {
			kind := "md"
			icon := assets.URL("/icons/file.png")
			if child.Type.IsDir() {
				kind = "dir"
				icon = assets.URL("/icons/folder.png")
			} else if strings.HasSuffix(strings.ToLower(paths[i]), ".sh") {
				// an executable opens a terminal, so it wears the terminal's icon here too
				kind = "script"
				icon = assets.URL("/icons/iterm.png")
			} else if child.Meta.ThumbPath != "" {
				icon = child.Meta.ThumbPath
			} else if child.Meta.ImagePath != "" {
				icon = child.Meta.ImagePath
			}
			entries = append(entries, EntryVM{
				Name:           path.Base(paths[i]),
				Path:           paths[i],
				Kind:           kind,
				Icon:           icon,
				ImagePath:      child.Meta.ImagePath,
				CompressedPath: child.Meta.CompressedPath,
				Size:           child.Meta.Size,
				Modified:       child.Meta.Modified,
			})
		}
		vm.Content = ContentVM{Kind: "folder", Entries: entries}
	} else {
		vm.Content = ContentVM{Kind: "markdown", HTML: markdown.Render(node.Content)}
	}

	b, _ := json.Marshal(vm)
	return b
}

// The sidebar offers only places that exist. It used to list Recents, Shared, iCloud and Network,
// none of which is in this filesystem, so every one of them led to an empty window.
func buildSidebar(currentPath string, filesystem fs.FileSystem) []SidebarSectionVM {
	folder := assets.URL("/icons/folder.png")
	item := func(name, p string) SidebarItemVM {
		return SidebarItemVM{Name: name, Path: p, Icon: folder, Active: currentPath == p}
	}

	var favorites []SidebarItemVM
	for _, name := range []string{"Applications", "Desktop", "Documents", "Downloads", "Movies", "Music", "Pictures", "Public"} {
		if filesystem.Exists("~/" + name) {
			favorites = append(favorites, item(name, "~/"+name))
		}
	}

	return []SidebarSectionVM{
		{Title: "Favorites", Items: favorites},
		{Title: "Locations", Items: []SidebarItemVM{item("farquaad", "~")}},
	}
}

func buildBreadcrumb(currentPath string) []BreadcrumbItemVM {
	parts := splitPath(currentPath)
	var items []BreadcrumbItemVM
	for _, p := range parts {
		items = append(items, BreadcrumbItemVM{
			Name: p.Name,
			Path: p.Path,
			Icon: assets.URL("/icons/folder.png"),
		})
	}
	return items
}

type pathPart struct {
	Name string
	Path string
}

func splitPath(p string) []pathPart {
	if p == "~" {
		return []pathPart{{Name: "farquaad", Path: "~"}}
	}
	// Strip ~ prefix and split
	clean := p
	if len(p) > 1 && p[0] == '~' {
		clean = p[2:] // skip "~/"
	}
	segments := []string{"Macintosh HD", "Users", "farquaad"}
	if clean != "" {
		segments = append(segments, splitCleanPath(clean)...)
	}
	var parts []pathPart
	full := ""
	for _, s := range segments {
		if full == "" {
			full = s
		} else {
			full = full + "/" + s
		}
		if s == "farquaad" {
			parts = append(parts, pathPart{Name: s, Path: "~"})
		} else {
			parts = append(parts, pathPart{Name: s, Path: full})
		}
	}
	return parts
}

func splitCleanPath(p string) []string {
	var parts []string
	start := 0
	for i := 0; i < len(p); i++ {
		if p[i] == '/' {
			if i > start {
				parts = append(parts, p[start:i])
			}
			start = i + 1
		}
	}
	if start < len(p) {
		parts = append(parts, p[start:])
	}
	return parts
}
