package wm

// Window is a generic OS window. App-specific state lives in separate per-app maps.
type Window struct {
	ID        string
	AppType   string // "finder" | "terminal" | "viewer" | "ephemeral" | "placeholder" | "editor"
	AppID     string // "finder" | "terminal" | "obsidian" | "zed" | etc.
	Title     string
	X, Y      int
	Width     int
	Height    int
	Z         int
	Minimized bool
	Maximized bool
	Ephemeral bool // no titlebar, no resize, click-away closes
}
