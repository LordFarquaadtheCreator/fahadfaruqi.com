package wm

import "encoding/json"

// WindowVM is the JSON-serializable view model for one window.
type WindowVM struct {
	ID        string          `json:"id"`
	AppType   string          `json:"appType"`
	AppID     string          `json:"appId"`
	Title     string          `json:"title"`
	X         int             `json:"x"`
	Y         int             `json:"y"`
	Z         int             `json:"z"`
	Width     int             `json:"width"`
	Height    int             `json:"height"`
	Minimized bool            `json:"minimized"`
	Maximized bool            `json:"maximized"`
	Content   json.RawMessage `json:"content"`
}

// BuildWindowsVM builds window VMs in z-order. Content is filled by app-specific builders.
type ContentBuilder func(win *Window) json.RawMessage

func BuildWindowsVM(m *Manager, contentBuilder ContentBuilder) []WindowVM {
	m.AssignZ()
	var windows []WindowVM
	for _, id := range m.order {
		w := m.windows[id]
		if w.Minimized {
			continue
		}
		var content json.RawMessage
		if contentBuilder != nil {
			content = contentBuilder(w)
		}
		windows = append(windows, WindowVM{
			ID:        w.ID,
			AppType:   w.AppType,
			AppID:     w.AppID,
			Title:     w.Title,
			X:         w.X,
			Y:         w.Y,
			Z:         w.Z,
			Width:     w.Width,
			Height:    w.Height,
			Minimized: w.Minimized,
			Maximized: w.Maximized,
			Content:   content,
		})
	}
	return windows
}
