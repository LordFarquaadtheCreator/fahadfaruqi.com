package dock

import (
	"encoding/json"
	"interpreter/apps"
	"interpreter/assets"
	"interpreter/wm"
)

// DockVM is the dock view model.
type DockVM struct {
	Items []DockItemVM `json:"items"`
	Trash DockItemVM   `json:"trash"`
}

type DockItemVM struct {
	AppID   string `json:"appId"`
	Name    string `json:"name"`
	Icon    string `json:"icon"`
	Running bool   `json:"running"`
	Active  bool   `json:"active"`
}

// BuildVM builds the dock view model from the app registry and window manager.
func BuildVM(m *wm.Manager) DockVM {
	items := make([]DockItemVM, 0, len(apps.DockOrder))

	// Track which apps have running windows
	runningApps := make(map[string]bool)
	activeAppID := ""
	focusedID := m.FocusedID()
	if focusedID != "" {
		if win, ok := m.Get(focusedID); ok {
			activeAppID = win.AppID
		}
	}
	for _, win := range m.Windows() {
		runningApps[win.AppID] = true
	}

	listed := make(map[string]bool, len(apps.DockOrder))
	for _, appID := range apps.DockOrder {
		app, ok := apps.Get(appID)
		if !ok {
			continue
		}
		listed[app.ID] = true
		items = append(items, DockItemVM{
			AppID:   app.ID,
			Name:    app.Name,
			Icon:    app.Icon,
			Running: runningApps[app.ID],
			Active:  app.ID == activeAppID,
		})
	}

	// A dock's job is to show what is open, so an app the visitor launched that the fixed order does
	// not list — the rose, a script's terminal — joins it, in the order its windows were opened.
	for _, win := range m.Windows() {
		if listed[win.AppID] {
			continue
		}
		app, ok := apps.Get(win.AppID)
		if !ok {
			continue
		}
		listed[win.AppID] = true
		items = append(items, DockItemVM{
			AppID:   app.ID,
			Name:    app.Name,
			Icon:    app.Icon,
			Running: true,
			Active:  app.ID == activeAppID,
		})
	}

	return DockVM{
		Items: items,
		Trash: DockItemVM{
			AppID: "trash",
			Name:  "Trash",
			Icon:  assets.URL("/icons/trash.png"),
		},
	}
}

// JSON marshals DockVM to JSON.
func (vm DockVM) JSON() string {
	b, err := json.Marshal(vm)
	if err != nil {
		return `{"items":[]}`
	}
	return string(b)
}
