package world

import (
	"interpreter/desktop"
	"interpreter/dock"
	"interpreter/finder"
	"interpreter/fs"
	"interpreter/menubar"
	"interpreter/terminal"
	"interpreter/viewer"
	"interpreter/wm"
)

type World struct {
	FS       fs.FileSystem
	Env      map[string]string
	WM       *wm.Manager
	Finder   *finder.FinderManager
	Terminal *terminal.TerminalManager
	Viewer   *viewer.ViewerManager
	Dock     *dock.DockVM
	MenuBar  *menubar.MenuBarVM
	Desktop  *desktop.DesktopVM
}
