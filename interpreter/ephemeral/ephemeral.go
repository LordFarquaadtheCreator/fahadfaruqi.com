package ephemeral

import (
	"encoding/json"
	"interpreter/apps"
	"interpreter/fs"
	"interpreter/wm"
	"math"
	"path"
	"strings"
)

type EphemeralState struct {
	FilePath  string
	FileName  string
	ImagePath string
	IsDir     bool
}

type EphemeralManager struct {
	wm         *wm.Manager
	states     map[string]*EphemeralState
	filesystem fs.FileSystem
	currentID  string
}

func NewEphemeralManager(w *wm.Manager, filesystem fs.FileSystem) *EphemeralManager {
	return &EphemeralManager{
		wm:         w,
		states:     make(map[string]*EphemeralState),
		filesystem: filesystem,
	}
}

func (em *EphemeralManager) SpawnForFile(filePath string) *wm.Window {
	// Close existing ephemeral if open
	em.Close()

	node, _ := em.filesystem.Stat(filePath)
	state := &EphemeralState{
		FilePath:  filePath,
		FileName:  path.Base(filePath),
		IsDir:     node.Type.IsDir(),
	}

	app, _ := apps.Get("preview-ephemeral")
	winW, winH := app.Width, app.Height

	if !state.IsDir && node.Meta.ImgWidth > 0 && node.Meta.ImgHeight > 0 {
		vw, vh := em.wm.Viewport()
		if vw > 0 && vh > 0 {
			imgW := float64(node.Meta.ImgWidth)
			imgH := float64(node.Meta.ImgHeight)

			scaleMin := math.Max(0.3*float64(vw)/imgW, 0.3*float64(vh)/imgH)
			scaleMax := math.Min(0.45*float64(vw)/imgW, 0.45*float64(vh)/imgH)
			naturalFit := math.Min(0.375*float64(vw)/imgW, 0.375*float64(vh)/imgH)

			scale := naturalFit
			if scale > scaleMax {
				scale = scaleMax
			}
			if scale < scaleMin {
				scale = scaleMin
			}

			winW = int(imgW*scale + 0.5)
			winH = int(imgH*scale + 0.5)
		}
		state.ImagePath = node.Meta.ImagePath
	}

	win := em.wm.SpawnCentered(string(app.Type), app.ID, state.FileName, winW, winH)
	em.states[win.ID] = state
	em.currentID = win.ID
	return win
}

func (em *EphemeralManager) Close() {
	if em.currentID == "" {
		return
	}
	em.wm.Close(em.currentID)
	delete(em.states, em.currentID)
	em.currentID = ""
}

func (em *EphemeralManager) CurrentID() string {
	return em.currentID
}

func (em *EphemeralManager) CurrentFilePath() string {
	if em.currentID == "" {
		return ""
	}
	s, ok := em.states[em.currentID]
	if !ok {
		return ""
	}
	return s.FilePath
}

func (em *EphemeralManager) Cleanup(id string) {
	delete(em.states, id)
	if em.currentID == id {
		em.currentID = ""
	}
}

func (em *EphemeralManager) GetState(id string) (*EphemeralState, bool) {
	s, ok := em.states[id]
	return s, ok
}

type EphemeralContentVM struct {
	Kind     string `json:"kind"` // "ephemeral-image" | "ephemeral-folder"
	Src      string `json:"src,omitempty"`
	Name     string `json:"name"`
	FilePath string `json:"filePath"`
	Icon     string `json:"icon,omitempty"`
}

func BuildContentVM(win *wm.Window, em *EphemeralManager) json.RawMessage {
	s, ok := em.GetState(win.ID)
	if !ok {
		return json.RawMessage(`{"kind":"ephemeral-folder","name":"","filePath":""}`)
	}

	vm := EphemeralContentVM{
		Name:     s.FileName,
		FilePath: s.FilePath,
	}

	if s.IsDir {
		vm.Kind = "ephemeral-folder"
		vm.Icon = "/icons/folder.png"
	} else {
		vm.Kind = "ephemeral-image"
		vm.Src = s.ImagePath
	}

	b, _ := json.Marshal(vm)
	return b
}

// IsImageFile checks if a file path has an image extension.
func IsImageFile(filePath string) bool {
	ext := strings.ToLower(path.Ext(filePath))
	switch ext {
	case ".png", ".jpg", ".jpeg", ".gif", ".webp":
		return true
	}
	return false
}
