package finder

import (
	"fmt"
	"interpreter/apps"
	"interpreter/fs"
	"interpreter/wm"
	"path"
)

type ViewMode string

const (
	IconView    ViewMode = "icon"
	ListView    ViewMode = "list"
	ColumnView  ViewMode = "column"
	GalleryView ViewMode = "gallery"
)

// FinderState holds per-window Finder navigation state.
type FinderState struct {
	Path          string
	ViewMode      ViewMode
	History       []string
	HistoryIndex  int
	SelectedEntry string
}

func (s *FinderState) CanBack() bool    { return s.HistoryIndex > 0 }
func (s *FinderState) CanForward() bool { return s.HistoryIndex < len(s.History)-1 }

func (s *FinderState) PushHistory(p string) {
	s.History = s.History[:s.HistoryIndex+1]
	if len(s.History) > 0 && s.History[len(s.History)-1] == p {
		return
	}
	s.History = append(s.History, p)
	s.HistoryIndex = len(s.History) - 1
}

func (s *FinderState) GoBack() {
	if s.CanBack() {
		s.HistoryIndex--
		s.Path = s.History[s.HistoryIndex]
	}
}

func (s *FinderState) GoForward() {
	if s.CanForward() {
		s.HistoryIndex++
		s.Path = s.History[s.HistoryIndex]
	}
}

// FinderManager owns per-window Finder state and delegates window ops to wm.Manager.
type FinderManager struct {
	wm      *wm.Manager
	states  map[string]*FinderState
	filesys fs.FileSystem
}

func NewFinderManager(w *wm.Manager, filesystem fs.FileSystem) *FinderManager {
	return &FinderManager{
		wm:      w,
		states:  make(map[string]*FinderState),
		filesys: filesystem,
	}
}

func (fm *FinderManager) Spawn(initialPath string) (*wm.Window, error) {
	if !fm.filesys.Exists(initialPath) {
		return nil, fmt.Errorf("path does not exist: %s", initialPath)
	}
	app, _ := apps.Get("finder")
	win := fm.wm.SpawnSilent(string(apps.TypeFinder), app.ID, path.Base(initialPath), app.Width, app.Height)
	fm.states[win.ID] = &FinderState{
		Path:     initialPath,
		ViewMode: IconView,
		History:  []string{initialPath},
	}
	return win, nil
}

func (fm *FinderManager) Navigate(id, destPath string) error {
	if !fm.filesys.Exists(destPath) {
		return fmt.Errorf("path does not exist: %s", destPath)
	}
	s, ok := fm.states[id]
	if !ok {
		return fmt.Errorf("window not found: %s", id)
	}
	s.Path = destPath
	s.PushHistory(destPath)
	s.SelectedEntry = ""
	if win, ok := fm.wm.Get(id); ok {
		win.Title = path.Base(destPath)
	}
	return nil
}

func (fm *FinderManager) Back(id string) {
	s, ok := fm.states[id]
	if !ok {
		return
	}
	s.GoBack()
	s.SelectedEntry = ""
	if win, ok := fm.wm.Get(id); ok {
		win.Title = path.Base(s.Path)
	}
}

func (fm *FinderManager) Forward(id string) {
	s, ok := fm.states[id]
	if !ok {
		return
	}
	s.GoForward()
	s.SelectedEntry = ""
	if win, ok := fm.wm.Get(id); ok {
		win.Title = path.Base(s.Path)
	}
}

func (fm *FinderManager) SetViewMode(id string, mode ViewMode) {
	s, ok := fm.states[id]
	if !ok {
		return
	}
	s.ViewMode = mode
}

func (fm *FinderManager) SelectEntry(id, name string) {
	s, ok := fm.states[id]
	if !ok {
		return
	}
	s.SelectedEntry = name
}

func (fm *FinderManager) GetState(id string) (*FinderState, bool) {
	s, ok := fm.states[id]
	return s, ok
}

func (fm *FinderManager) Cleanup(id string) {
	delete(fm.states, id)
}
