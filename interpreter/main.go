//go:build js && wasm

package main

import (
	"encoding/json"
	"interpreter/apps"
	"interpreter/desktop"
	"interpreter/dock"
	"interpreter/finder"
	"interpreter/menubar"
	"interpreter/session"
	"interpreter/shell"
	"interpreter/terminal"
	"interpreter/viewer"
	"interpreter/wm"
	"interpreter/world"
	"path"
	"syscall/js"
)

var (
	w      world.World
	s      *session.Session
	vmSubs []js.Value
)

// RootVM is the full desktop view model pushed to JS on every state change.
type RootVM struct {
	Windows   []wm.WindowVM    `json:"windows"`
	FocusedID string           `json:"focusedId"`
	ViewportW int              `json:"viewportW"`
	ViewportH int              `json:"viewportH"`
	Dock      dock.DockVM      `json:"dock"`
	MenuBar   menubar.MenuBarVM `json:"menuBar"`
	Desktop   desktop.DesktopVM `json:"desktop"`
}

func main() {
	w = world.Load()
	s = session.New(w.Env)
	w.WM.SetNotify(pushVM)

	// Init terminal with command runner adapter (breaks import cycle)
	w.Terminal = terminal.NewTerminalManager(w.WM, &terminalRunner{}, w.Env)

	js.Global().Set("interpRun", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		input := args[0].String()
		return shell.Run(input, w, s)
	}))

	os := make(map[string]interface{})

	os["subscribe"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		cb := args[0]
		vmSubs = append(vmSubs, cb)
		pushVM()
		return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			for i, sub := range vmSubs {
				if sub.Equal(cb) {
					vmSubs = append(vmSubs[:i], vmSubs[i+1:]...)
					break
				}
			}
			return nil
		})
	})

	os["setViewport"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.WM.SetViewport(args[0].Int(), args[1].Int())
		return nil
	})

	// Window operations
	os["spawn"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		appID := args[0].String()
		app, ok := apps.Get(appID)
		if !ok {
			return "unknown app: " + appID
		}
		switch app.Type {
		case apps.TypeFinder:
			w.Finder.Spawn("~")
		case apps.TypeTerminal:
			w.Terminal.Spawn()
		case apps.TypePlaceholder, apps.TypeEditor:
			w.WM.Spawn(string(app.Type), app.ID, app.Name, app.Width, app.Height)
		default:
			w.WM.Spawn(string(app.Type), app.ID, app.Name, app.Width, app.Height)
		}
		pushVM()
		return nil
	})

	os["openFile"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		filePath := args[0].String()
		if !w.FS.Exists(filePath) {
			return "path does not exist: " + filePath
		}
		if w.FS.IsDir(filePath) {
			w.Finder.Spawn(filePath)
			pushVM()
			return nil
		}
		appID := apps.AppForFile(path.Base(filePath))
		if appID == "viewer-image" || appID == "viewer-text" {
			w.Viewer.SpawnForFile(filePath)
		} else if appID != "" {
			app, _ := apps.Get(appID)
			w.WM.Spawn(string(app.Type), app.ID, path.Base(filePath), app.Width, app.Height)
		}
		pushVM()
		return nil
	})

	os["close"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		id := args[0].String()
		cleanupWindow(id)
		w.WM.Close(id)
		return nil
	})

	os["closeFocused"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		id := w.WM.FocusedID()
		if id != "" {
			cleanupWindow(id)
		}
		w.WM.CloseFocused()
		return nil
	})

	os["focus"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.WM.Focus(args[0].String())
		return nil
	})

	os["drag"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.WM.Drag(args[0].String(), args[1].Int(), args[2].Int())
		return nil
	})

	os["resize"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.WM.Resize(args[0].String(), args[1].Int(), args[2].Int())
		return nil
	})

	os["minimize"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.WM.Minimize(args[0].String())
		return nil
	})

	os["maximize"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.WM.Maximize(args[0].String())
		return nil
	})

	os["restore"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.WM.Restore(args[0].String())
		return nil
	})

	// Finder operations
	os["finderNavigate"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		err := w.Finder.Navigate(args[0].String(), args[1].String())
		if err != nil {
			return err.Error()
		}
		return nil
	})

	os["finderBack"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.Finder.Back(args[0].String())
		return nil
	})

	os["finderForward"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.Finder.Forward(args[0].String())
		return nil
	})

	os["finderSetView"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.Finder.SetViewMode(args[0].String(), finder.ViewMode(args[1].String()))
		return nil
	})

	os["finderSelect"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.Finder.SelectEntry(args[0].String(), args[1].String())
		return nil
	})

	os["finderOpenEntry"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		winID := args[0].String()
		entryPath := args[1].String()
		if w.FS.IsDir(entryPath) {
			if err := w.Finder.Navigate(winID, entryPath); err != nil {
				return err.Error()
			}
			pushVM()
			return nil
		}
		// File — open in viewer
		appID := apps.AppForFile(path.Base(entryPath))
		if appID == "viewer-image" || appID == "viewer-text" {
			w.Viewer.SpawnForFile(entryPath)
		} else if appID != "" {
			app, _ := apps.Get(appID)
			w.WM.Spawn(string(app.Type), app.ID, path.Base(entryPath), app.Width, app.Height)
		}
		pushVM()
		return nil
	})

	// Terminal operations
	os["terminalRun"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		output := w.Terminal.RunCommand(args[0].String(), args[1].String())
		pushVM()
		return output
	})

	os["terminalComplete"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		candidates := w.Terminal.TabComplete(args[0].String(), args[1].String())
		b, _ := json.Marshal(candidates)
		return string(b)
	})

	// Dock
	os["dockClick"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		appID := args[0].String()
		app, ok := apps.Get(appID)
		if !ok {
			return nil
		}
		// Check if app already has a window — focus it, else spawn
		for _, win := range w.WM.Windows() {
			if win.AppID == appID {
				w.WM.Focus(win.ID)
				return nil
			}
		}
		switch app.Type {
		case apps.TypeFinder:
			w.Finder.Spawn("~")
		case apps.TypeTerminal:
			w.Terminal.Spawn()
		default:
			w.WM.Spawn(string(app.Type), app.ID, app.Name, app.Width, app.Height)
		}
		pushVM()
		return nil
	})

	js.Global().Set("os", js.ValueOf(os))

	select {}
}

// cleanupWindow removes app-specific state when a window closes.
func cleanupWindow(id string) {
	w.Finder.Cleanup(id)
	w.Terminal.Cleanup(id)
	w.Viewer.Cleanup(id)
}

// buildContent dispatches to the right app-specific content builder.
func buildContent(win *wm.Window) json.RawMessage {
	switch win.AppType {
	case string(apps.TypeFinder):
		return finder.BuildContentVM(win, w.Finder, w.FS)
	case string(apps.TypeTerminal):
		return terminal.BuildContentVM(win, w.Terminal)
	case string(apps.TypeViewer):
		return viewer.BuildContentVM(win, w.Viewer)
	default:
		// placeholder, editor
		return buildSplashContent(win)
	}
}

type splashContentVM struct {
	Kind  string `json:"kind"` // "splash"
	AppID string `json:"appId"`
	Name  string `json:"name"`
	Icon  string `json:"icon"`
}

func buildSplashContent(win *wm.Window) json.RawMessage {
	app, ok := apps.Get(win.AppID)
	if !ok {
		return json.RawMessage(`{"kind":"splash","appId":"","name":"Unknown","icon":""}`)
	}
	vm := splashContentVM{
		Kind:  "splash",
		AppID: app.ID,
		Name:  app.Name,
		Icon:  app.Icon,
	}
	b, _ := json.Marshal(vm)
	return b
}

func pushVM() {
	vpW, vpH := w.WM.Viewport()
	root := RootVM{
		Windows:   wm.BuildWindowsVM(w.WM, buildContent),
		FocusedID: w.WM.FocusedID(),
		ViewportW: vpW,
		ViewportH: vpH,
		Dock:      dock.BuildVM(w.WM),
		MenuBar:   menubar.BuildVM(w.WM),
		Desktop:   desktop.BuildVM(w.FS),
	}
	b, err := json.Marshal(root)
	if err != nil {
		return
	}
	jsonStr := string(b)
	for _, cb := range vmSubs {
		cb.Invoke(jsonStr)
	}
}

// terminalRunner implements terminal.CommandRunner by delegating to shell.Run and terminal.Complete.
type terminalRunner struct{}

func (tr *terminalRunner) RunCommand(input string, s *session.Session) string {
	return shell.Run(input, w, s)
}

func (tr *terminalRunner) TabComplete(input, cwd string) []string {
	return terminal.Complete(input, cwd, w.FS)
}
