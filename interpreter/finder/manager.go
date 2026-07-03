package finder

import (
	"fmt"
	"interpreter/fs"
)

const (
	defaultWidth  = 640
	defaultHeight = 480
	maxWindows    = 10
)

// Manager owns all Finder window state.
type Manager struct {
	windows   map[string]*Window
	order     []string // z-order; last element is focused/topmost
	nextID    int
	viewportW int
	viewportH int
	notify    func()
}

func NewManager() *Manager {
	return &Manager{
		windows: make(map[string]*Window),
		order:   make([]string, 0),
	}
}

func (m *Manager) SetNotify(n func()) { m.notify = n }

func (m *Manager) doNotify() {
	if m.notify != nil {
		m.notify()
	}
}

func (m *Manager) SetViewport(w, h int) {
	m.viewportW = w
	m.viewportH = h
	m.clampAll()
	m.doNotify()
}

func (m *Manager) Spawn(path string, filesystem fs.FileSystem) (*Window, error) {
	if !filesystem.Exists(path) {
		return nil, fmt.Errorf("path does not exist: %s", path)
	}
	if len(m.windows) >= maxWindows && len(m.order) > 0 {
		m.closeWindow(m.order[0])
	}
	id := fmt.Sprintf("win-%d", m.nextID)
	m.nextID++
	win := &Window{
		ID:       id,
		Path:     path,
		ViewMode: IconView,
		Width:    defaultWidth,
		Height:   defaultHeight,
		History:  []string{path},
	}
	offset := (len(m.windows) % 5) * 30
	win.X = 50 + offset
	win.Y = 50 + offset
	m.clampWindow(win)
	m.windows[id] = win
	m.order = append(m.order, id)
	m.doNotify()
	return win, nil
}

func (m *Manager) Close(id string) {
	m.closeWindow(id)
	m.doNotify()
}

func (m *Manager) CloseFocused() {
	if len(m.order) == 0 {
		return
	}
	m.closeWindow(m.order[len(m.order)-1])
	m.doNotify()
}

func (m *Manager) closeWindow(id string) {
	delete(m.windows, id)
	newOrder := make([]string, 0, len(m.order))
	for _, o := range m.order {
		if o != id {
			newOrder = append(newOrder, o)
		}
	}
	m.order = newOrder
}

func (m *Manager) Focus(id string) {
	newOrder := make([]string, 0, len(m.order))
	for _, o := range m.order {
		if o != id {
			newOrder = append(newOrder, o)
		}
	}
	newOrder = append(newOrder, id)
	m.order = newOrder
	m.doNotify()
}

func (m *Manager) Drag(id string, dx, dy int) {
	w, ok := m.windows[id]
	if !ok {
		return
	}
	w.X += dx
	w.Y += dy
	m.clampWindow(w)
	m.doNotify()
}

func (m *Manager) Navigate(id, path string, filesystem fs.FileSystem) error {
	if !filesystem.Exists(path) {
		return fmt.Errorf("path does not exist: %s", path)
	}
	w, ok := m.windows[id]
	if !ok {
		return fmt.Errorf("window not found: %s", id)
	}
	w.Path = path
	w.PushHistory(path)
	w.SelectedEntry = ""
	m.doNotify()
	return nil
}

func (m *Manager) Back(id string) {
	w, ok := m.windows[id]
	if !ok {
		return
	}
	w.GoBack()
	w.SelectedEntry = ""
	m.doNotify()
}

func (m *Manager) Forward(id string) {
	w, ok := m.windows[id]
	if !ok {
		return
	}
	w.GoForward()
	w.SelectedEntry = ""
	m.doNotify()
}

func (m *Manager) SetViewMode(id string, mode ViewMode) {
	w, ok := m.windows[id]
	if !ok {
		return
	}
	w.ViewMode = mode
	m.doNotify()
}

func (m *Manager) SelectEntry(id, name string) {
	w, ok := m.windows[id]
	if !ok {
		return
	}
	w.SelectedEntry = name
	m.doNotify()
}

// OpenEntry handles double-click. Dirs navigate in-place; files spawn a new window.
func (m *Manager) OpenEntry(id, path string, filesystem fs.FileSystem) error {
	if filesystem.IsDir(path) {
		return m.Navigate(id, path, filesystem)
	}
	_, err := m.Spawn(path, filesystem)
	return err
}

func (m *Manager) clampWindow(w *Window) {
	minVisible := 40
	if m.viewportW > 0 {
		if w.X+w.Width < minVisible {
			w.X = minVisible - w.Width
		}
		if w.X > m.viewportW-minVisible {
			w.X = m.viewportW - minVisible
		}
		if w.Y < 0 {
			w.Y = 0
		}
		if w.Y > m.viewportH-minVisible {
			w.Y = m.viewportH - minVisible
		}
	}
}

func (m *Manager) clampAll() {
	for _, w := range m.windows {
		m.clampWindow(w)
	}
}

func (m *Manager) assignZ() {
	for i, id := range m.order {
		if w, ok := m.windows[id]; ok {
			w.Z = i + 1
		}
	}
}
