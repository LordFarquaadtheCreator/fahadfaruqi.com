package finder

import (
	"encoding/json"
	"interpreter/fs"
	"path"
)

// VM is the root ViewModel pushed to JS on every state change.
type VM struct {
	Windows   []WindowVM `json:"windows"`
	FocusedID string     `json:"focusedId"`
	ViewportW int        `json:"viewportW"`
	ViewportH int        `json:"viewportH"`
}

// WindowVM describes one Finder window.
type WindowVM struct {
	ID            string          `json:"id"`
	Title         string          `json:"title"`
	X             int             `json:"x"`
	Y             int             `json:"y"`
	Z             int             `json:"z"`
	Width         int             `json:"width"`
	Height        int             `json:"height"`
	ViewMode      string          `json:"viewMode"`
	SelectedEntry string          `json:"selectedEntry"`
	CanBack       bool            `json:"canBack"`
	CanForward    bool            `json:"canForward"`
	Sidebar       []SidebarItemVM `json:"sidebar"`
	Content       ContentVM       `json:"content"`
}

// SidebarItemVM is one row in the sidebar tree.
type SidebarItemVM struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Kind   string `json:"kind"` // "root" | "dir"
	Level  int    `json:"level"`
	Active bool   `json:"active"`
}

// ContentVM describes what is shown in the main pane.
type ContentVM struct {
	Kind    string    `json:"kind"`    // "folder" | "markdown"
	Entries []EntryVM `json:"entries,omitempty"`
	HTML    string    `json:"html,omitempty"`
}

// EntryVM is a file or folder inside a directory listing.
type EntryVM struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Kind     string `json:"kind"` // "dir" | "md"
	Size     int64  `json:"size"`
	Modified string `json:"modified"`
}

// BuildVM builds the full ViewModel from Manager + FileSystem.
func BuildVM(m *Manager, filesystem fs.FileSystem) VM {
	m.assignZ()
	var windows []WindowVM
	focusedID := ""
	if len(m.order) > 0 {
		focusedID = m.order[len(m.order)-1]
	}
	for _, id := range m.order {
		w := m.windows[id]
		windows = append(windows, buildWindowVM(w, filesystem))
	}
	return VM{
		Windows:   windows,
		FocusedID: focusedID,
		ViewportW: m.viewportW,
		ViewportH: m.viewportH,
	}
}

func buildWindowVM(w *Window, filesystem fs.FileSystem) WindowVM {
	vm := WindowVM{
		ID:            w.ID,
		Title:         w.Title(),
		X:             w.X,
		Y:             w.Y,
		Z:             w.Z,
		Width:         w.Width,
		Height:        w.Height,
		ViewMode:      string(w.ViewMode),
		SelectedEntry: w.SelectedEntry,
		CanBack:       w.CanBack(),
		CanForward:    w.CanForward(),
		Sidebar:       buildSidebar(w.Path, filesystem),
	}
	node, err := filesystem.Stat(w.Path)
	if err != nil {
		vm.Content = ContentVM{Kind: "folder", Entries: []EntryVM{}}
		return vm
	}
	if node.Type.IsDir() {
		children, paths, _ := filesystem.Children(w.Path, false)
		var entries []EntryVM
		for i, child := range children {
			kind := "md"
			if child.Type.IsDir() {
				kind = "dir"
			}
			entries = append(entries, EntryVM{
				Name:     path.Base(paths[i]),
				Path:     paths[i],
				Kind:     kind,
				Size:     child.Meta.Size,
				Modified: child.Meta.Modified,
			})
		}
		vm.Content = ContentVM{Kind: "folder", Entries: entries}
	} else {
		vm.Content = ContentVM{Kind: "markdown", HTML: RenderMarkdown(node.Content)}
	}
	return vm
}

func buildSidebar(currentPath string, filesystem fs.FileSystem) []SidebarItemVM {
	items := []SidebarItemVM{
		{Name: "Home", Path: "~", Kind: "root", Level: 0, Active: currentPath == "~"},
	}
	items = append(items, buildDirTree("~", currentPath, filesystem, 1)...)
	return items
}

func buildDirTree(dirPath, currentPath string, filesystem fs.FileSystem, level int) []SidebarItemVM {
	var items []SidebarItemVM
	children, paths, _ := filesystem.Children(dirPath, false)
	for i, child := range children {
		if !child.Type.IsDir() {
			continue
		}
		items = append(items, SidebarItemVM{
			Name:   path.Base(paths[i]),
			Path:   paths[i],
			Kind:   "dir",
			Level:  level,
			Active: paths[i] == currentPath,
		})
		items = append(items, buildDirTree(paths[i], currentPath, filesystem, level+1)...)
	}
	return items
}

// JSON marshals VM to JSON.
func (vm VM) JSON() string {
	b, err := json.Marshal(vm)
	if err != nil {
		return "{}"
	}
	return string(b)
}
