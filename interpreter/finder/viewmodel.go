package finder

import (
	"encoding/json"
	"interpreter/fs"
	"interpreter/wm"
	"path"
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
	Title string           `json:"title"` // "", "Favorites", "Locations"
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
	Name     string `json:"name"`
	Path     string `json:"path"`
	Kind     string `json:"kind"` // "dir" | "md" | "txt" | "png" | "jpeg"
	Icon     string `json:"icon"`
	Size     int64  `json:"size"`
	Modified string `json:"modified"`
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
		Sidebar:       buildSidebar(s.Path),
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
			icon := "/icons/file.png"
			if child.Type.IsDir() {
				kind = "dir"
				icon = "/icons/folder.png"
			}
			entries = append(entries, EntryVM{
				Name:     path.Base(paths[i]),
				Path:     paths[i],
				Kind:     kind,
				Icon:     icon,
				Size:     child.Meta.Size,
				Modified: child.Meta.Modified,
			})
		}
		vm.Content = ContentVM{Kind: "folder", Entries: entries}
	} else {
		vm.Content = ContentVM{Kind: "markdown", HTML: RenderMarkdown(node.Content)}
	}

	b, _ := json.Marshal(vm)
	return b
}

func buildSidebar(currentPath string) []SidebarSectionVM {
	return []SidebarSectionVM{
		{
			Title: "",
			Items: []SidebarItemVM{
				{Name: "Recents", Path: "~/Recents", Icon: "/icons/folder.png", Active: currentPath == "~/Recents"},
				{Name: "Shared", Path: "~/Shared", Icon: "/icons/folder.png", Active: currentPath == "~/Shared"},
			},
		},
		{
			Title: "Favorites",
			Items: []SidebarItemVM{
				{Name: "Applications", Path: "~/Applications", Icon: "/icons/folder.png", Active: currentPath == "~/Applications"},
				{Name: "Desktop", Path: "~/Desktop", Icon: "/icons/folder.png", Active: currentPath == "~/Desktop"},
				{Name: "Documents", Path: "~/Documents", Icon: "/icons/folder.png", Active: currentPath == "~/Documents"},
				{Name: "Downloads", Path: "~/Downloads", Icon: "/icons/folder.png", Active: currentPath == "~/Downloads"},
			},
		},
		{
			Title: "Locations",
			Items: []SidebarItemVM{
				{Name: "iCloud Drive", Path: "~/iCloud", Icon: "/icons/folder.png", Active: currentPath == "~/iCloud"},
				{Name: "farquaad", Path: "~", Icon: "/icons/folder.png", Active: currentPath == "~"},
				{Name: "Network", Path: "~/Network", Icon: "/icons/folder.png", Active: currentPath == "~/Network"},
			},
		},
	}
}

func buildBreadcrumb(currentPath string) []BreadcrumbItemVM {
	parts := splitPath(currentPath)
	var items []BreadcrumbItemVM
	for _, p := range parts {
		items = append(items, BreadcrumbItemVM{
			Name: p.Name,
			Path: p.Path,
			Icon: "/icons/folder.png",
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
