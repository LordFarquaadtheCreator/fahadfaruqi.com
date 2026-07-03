package dock

import (
	"encoding/json"
	"interpreter/apps"
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

	for _, appID := range apps.DockOrder {
		app, ok := apps.Get(appID)
		if !ok {
			continue
		}
		items = append(items, DockItemVM{
			AppID:   app.ID,
			Name:    app.Name,
			Icon:    app.Icon,
			Running: runningApps[app.ID],
			Active:  app.ID == activeAppID,
		})
	}

	return DockVM{
		Items: items,
		Trash: DockItemVM{
			AppID: "trash",
			Name:  "Trash",
			Icon:  "/icons/trash.png",
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
