package viewer

import (
	"encoding/json"
	"interpreter/apps"
	"interpreter/finder"
	"interpreter/fs"
	"interpreter/wm"
	"path"
	"strings"
)

// ViewerState holds per-window viewer state.
type ViewerState struct {
	FilePath string
	FileName string
	MimeType string
	Content  string
	ImagePath string // for image files — URL to serve
}

// ViewerManager owns per-window viewer state.
type ViewerManager struct {
	wm        *wm.Manager
	states    map[string]*ViewerState
	filesystem fs.FileSystem
}

func NewViewerManager(w *wm.Manager, filesystem fs.FileSystem) *ViewerManager {
	return &ViewerManager{
		wm:         w,
		states:     make(map[string]*ViewerState),
		filesystem: filesystem,
	}
}

// SpawnForFile creates a viewer window for the given file path.
func (vm *ViewerManager) SpawnForFile(filePath string) *wm.Window {
	appID := apps.AppForFile(filePath)
	if appID == "" {
		appID = "viewer-text"
	}

	app, _ := apps.Get(appID)
	node, _ := vm.filesystem.Stat(filePath)

	state := &ViewerState{
		FilePath: filePath,
		FileName: path.Base(filePath),
	}

	ext := strings.ToLower(path.Ext(filePath))
	switch ext {
	case ".png", ".jpg", ".jpeg", ".gif", ".webp":
		state.MimeType = "image/*"
		state.ImagePath = node.Meta.ImagePath
	case ".md":
		state.MimeType = "text/markdown"
		state.Content = finder.RenderMarkdown(node.Content)
	default:
		state.MimeType = "text/plain"
		state.Content = node.Content
	}

	win := vm.wm.Spawn(string(app.Type), app.ID, state.FileName, app.Width, app.Height)
	vm.states[win.ID] = state
	return win
}

func (vm *ViewerManager) GetState(id string) (*ViewerState, bool) {
	s, ok := vm.states[id]
	return s, ok
}

func (vm *ViewerManager) Cleanup(id string) {
	delete(vm.states, id)
}

// ViewerContentVM is the app-specific content for a viewer window.
type ViewerContentVM struct {
	Kind      string `json:"kind"` // "viewer-image" | "viewer-text"
	Src       string `json:"src,omitempty"`
	Name      string `json:"name"`
	Content   string `json:"content,omitempty"`
	IsMarkdown bool  `json:"isMarkdown,omitempty"`
}

// BuildContentVM builds viewer-specific content for a window.
func BuildContentVM(win *wm.Window, vm *ViewerManager) json.RawMessage {
	s, ok := vm.GetState(win.ID)
	if !ok {
		return json.RawMessage(`{"kind":"viewer-text","name":"","content":""}`)
	}

	var vcm ViewerContentVM
	if s.MimeType == "image/*" {
		vcm = ViewerContentVM{
			Kind: "viewer-image",
			Src:  s.ImagePath,
			Name: s.FileName,
		}
	} else {
		vcm = ViewerContentVM{
			Kind:       "viewer-text",
			Name:       s.FileName,
			Content:    s.Content,
			IsMarkdown: s.MimeType == "text/markdown",
		}
	}

	b, _ := json.Marshal(vcm)
	return b
}
