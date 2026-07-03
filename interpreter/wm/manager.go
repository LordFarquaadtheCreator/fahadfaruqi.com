package wm

import "fmt"

const maxWindows = 10

// Manager owns all generic window state: position, size, z-order, minimize/maximize.
// App-specific state (Finder navigation, Terminal history, etc.) lives in separate packages.
type Manager struct {
	windows   map[string]*Window
	order     []string // z-order; last element is focused/topmost
	nextID    int
	viewportW int
	viewportH int
	// lastDragX/lastDragY track position of most recently dragged window
	// so new windows spawn offset from it instead of screen center
	lastDragX int
	lastDragY int
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

func (m *Manager) Spawn(appType, appID, title string, width, height int) *Window {
	return m.spawn(appType, appID, title, width, height, true)
}

// SpawnSilent creates a window without triggering notify. Caller must call
// the app-specific state registration, then pushVM() manually.
func (m *Manager) SpawnSilent(appType, appID, title string, width, height int) *Window {
	return m.spawn(appType, appID, title, width, height, false)
}

func (m *Manager) spawn(appType, appID, title string, width, height int, notify bool) *Window {
	if len(m.windows) >= maxWindows && len(m.order) > 0 {
		m.closeWindow(m.order[0])
	}
	id := fmt.Sprintf("win-%d", m.nextID)
	m.nextID++
	win := &Window{
		ID:      id,
		AppType: appType,
		AppID:   appID,
		Title:   title,
		Width:   width,
		Height:  height,
	}
	offset := (len(m.windows) % 5) * 30
	if m.lastDragX != 0 || m.lastDragY != 0 {
		win.X = m.lastDragX + offset
		win.Y = m.lastDragY + offset
	} else if m.viewportW > 0 && m.viewportH > 0 {
		win.X = (m.viewportW-width)/2 + offset
		win.Y = (m.viewportH-height)/2 + offset
	} else {
		win.X = 50 + offset
		win.Y = 50 + offset
	}
	m.clampWindow(win)
	m.windows[id] = win
	m.order = append(m.order, id)
	if notify {
		m.doNotify()
	}
	return win
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
	if _, ok := m.windows[id]; !ok {
		return
	}
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
	m.lastDragX = w.X
	m.lastDragY = w.Y
	m.doNotify()
}

func (m *Manager) Resize(id string, width, height int) {
	w, ok := m.windows[id]
	if !ok {
		return
	}
	w.Width = width
	w.Height = height
	m.clampWindow(w)
	m.doNotify()
}

func (m *Manager) Minimize(id string) {
	w, ok := m.windows[id]
	if !ok {
		return
	}
	w.Minimized = true
	m.doNotify()
}

func (m *Manager) Maximize(id string) {
	w, ok := m.windows[id]
	if !ok {
		return
	}
	if w.Maximized {
		w.Maximized = false
	} else {
		w.Maximized = true
		w.Minimized = false
	}
	m.doNotify()
}

func (m *Manager) Restore(id string) {
	w, ok := m.windows[id]
	if !ok {
		return
	}
	w.Minimized = false
	m.doNotify()
}

func (m *Manager) Get(id string) (*Window, bool) {
	w, ok := m.windows[id]
	return w, ok
}

func (m *Manager) FocusedID() string {
	if len(m.order) == 0 {
		return ""
	}
	return m.order[len(m.order)-1]
}

func (m *Manager) Order() []string {
	return m.order
}

func (m *Manager) Windows() map[string]*Window {
	return m.windows
}

func (m *Manager) Viewport() (int, int) {
	return m.viewportW, m.viewportH
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

func (m *Manager) AssignZ() {
	for i, id := range m.order {
		if w, ok := m.windows[id]; ok {
			w.Z = i + 1
		}
	}
}
