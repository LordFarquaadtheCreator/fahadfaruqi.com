package world

import (
	"interpreter/desktop"
	"interpreter/dock"
	"interpreter/ephemeral"
	"interpreter/finder"
	"interpreter/fs"
	"interpreter/menubar"
	"interpreter/rose"
	"interpreter/terminal"
	"interpreter/viewer"
	"interpreter/wm"
)

type World struct {
	FS        fs.FileSystem
	Env       map[string]string
	WM        *wm.Manager
	Finder    *finder.FinderManager
	Terminal  *terminal.TerminalManager
	Viewer    *viewer.ViewerManager
	Ephemeral *ephemeral.EphemeralManager
	Rose      *rose.RoseManager
	Dock      *dock.DockVM
	MenuBar   *menubar.MenuBarVM
	Desktop   *desktop.DesktopVM
}
