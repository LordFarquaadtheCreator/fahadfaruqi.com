package menubar

import (
	"encoding/json"
	"interpreter/apps"
	"interpreter/wm"
	"time"
)

// MenuBarVM is the menu bar view model.
type MenuBarVM struct {
	ActiveApp string       `json:"activeApp"`
	Menus     []MenuVM     `json:"menus"`
	Clock     string       `json:"clock"`
	Tray      []TrayItemVM `json:"tray"`
}

type MenuVM struct {
	Name   string     `json:"name"`
	Items  []MenuItemVM `json:"items,omitempty"`
}

type MenuItemVM struct {
	Label    string `json:"label"`
	Shortcut string `json:"shortcut,omitempty"`
	Separator bool  `json:"separator,omitempty"`
}

type TrayItemVM struct {
	Icon string `json:"icon"`
	Name string `json:"name"` // tooltip label
}

// BuildVM builds the menu bar view model.
func BuildVM(m *wm.Manager) MenuBarVM {
	activeApp := "Finder"
	focusedID := m.FocusedID()
	if focusedID != "" {
		if win, ok := m.Get(focusedID); ok {
			if app, ok := apps.Get(win.AppID); ok {
				activeApp = app.Name
			}
		}
	}

	menus := buildMenus(activeApp)
	clock := time.Now().Format("Mon Jan 2 3:04 PM")
	tray := []TrayItemVM{
		{Icon: "search", Name: "Search"},
		{Icon: "controls", Name: "Control Center"},
		{Icon: "wifi", Name: "Wi-Fi"},
		{Icon: "battery-full", Name: "Battery"},
	}

	return MenuBarVM{
		ActiveApp: activeApp,
		Menus:     menus,
		Clock:     clock,
		Tray:      tray,
	}
}

func buildMenus(appName string) []MenuVM {
	return []MenuVM{
		{
			Name: "",
			Items: []MenuItemVM{
				{Label: "About This Mac"},
				{Separator: true},
				{Label: "System Settings..."},
				{Label: "App Store..."},
				{Separator: true},
				{Label: "Sleep"},
				{Label: "Restart..."},
				{Label: "Shut Down..."},
				{Separator: true},
				{Label: "Lock Screen", Shortcut: "Ctrl+Cmd+Q"},
				{Label: "Log Out", Shortcut: "Shift+Cmd+Q"},
			},
		},
		{
			Name: appName,
			Items: []MenuItemVM{
				{Label: "About " + appName},
				{Separator: true},
				{Label: "Settings...", Shortcut: "Cmd+,"},
				{Separator: true},
				{Label: "Hide " + appName, Shortcut: "Cmd+H"},
				{Label: "Hide Others", Shortcut: "Alt+Cmd+H"},
				{Separator: true},
				{Label: "Quit " + appName, Shortcut: "Cmd+Q"},
			},
		},
		{
			Name: "File",
			Items: []MenuItemVM{
				{Label: "New Window", Shortcut: "Cmd+N"},
				{Label: "New Tab", Shortcut: "Cmd+T"},
				{Separator: true},
				{Label: "Open...", Shortcut: "Cmd+O"},
				{Label: "Close Window", Shortcut: "Cmd+W"},
			},
		},
		{
			Name: "Edit",
			Items: []MenuItemVM{
				{Label: "Undo", Shortcut: "Cmd+Z"},
				{Label: "Redo", Shortcut: "Shift+Cmd+Z"},
				{Separator: true},
				{Label: "Cut", Shortcut: "Cmd+X"},
				{Label: "Copy", Shortcut: "Cmd+C"},
				{Label: "Paste", Shortcut: "Cmd+V"},
				{Label: "Select All", Shortcut: "Cmd+A"},
			},
		},
		{
			Name: "View",
			Items: []MenuItemVM{
				{Label: "as Icons", Shortcut: "Cmd+1"},
				{Label: "as List", Shortcut: "Cmd+2"},
				{Label: "as Columns", Shortcut: "Cmd+3"},
				{Label: "as Gallery", Shortcut: "Cmd+4"},
			},
		},
		{
			Name: "Go",
			Items: []MenuItemVM{
				{Label: "Back", Shortcut: "Cmd+["},
				{Label: "Forward", Shortcut: "Cmd+]"},
				{Separator: true},
				{Label: "Home", Shortcut: "Shift+Cmd+H"},
			},
		},
		{
			Name: "Window",
			Items: []MenuItemVM{
				{Label: "Minimize", Shortcut: "Cmd+M"},
				{Label: "Zoom"},
			},
		},
		{
			Name: "Help",
			Items: []MenuItemVM{
				{Label: appName + " Help"},
			},
		},
	}
}

// JSON marshals MenuBarVM to JSON.
func (vm MenuBarVM) JSON() string {
	b, err := json.Marshal(vm)
	if err != nil {
		return `{}`
	}
	return string(b)
}
