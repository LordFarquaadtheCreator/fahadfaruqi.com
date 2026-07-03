package finder

import "path"

type ViewMode string

const (
	IconView    ViewMode = "icon"
	ListView    ViewMode = "list"
	ColumnView  ViewMode = "column"
	GalleryView ViewMode = "gallery"
)

// Window represents a single Finder window.
type Window struct {
	ID            string
	Path          string
	ViewMode      ViewMode
	X, Y          int
	Width, Height int
	Z             int
	SelectedEntry string
	History       []string
	HistoryIndex  int
}

func (w *Window) Title() string { return path.Base(w.Path) }

func (w *Window) CanBack() bool    { return w.HistoryIndex > 0 }
func (w *Window) CanForward() bool { return w.HistoryIndex < len(w.History)-1 }

func (w *Window) PushHistory(p string) {
	w.History = w.History[:w.HistoryIndex+1]
	if len(w.History) > 0 && w.History[len(w.History)-1] == p {
		return
	}
	w.History = append(w.History, p)
	w.HistoryIndex = len(w.History) - 1
}

func (w *Window) GoBack() {
	if w.CanBack() {
		w.HistoryIndex--
		w.Path = w.History[w.HistoryIndex]
	}
}

func (w *Window) GoForward() {
	if w.CanForward() {
		w.HistoryIndex++
		w.Path = w.History[w.HistoryIndex]
	}
}
