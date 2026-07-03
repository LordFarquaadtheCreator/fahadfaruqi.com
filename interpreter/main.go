//go:build js && wasm

package main

import (
	"syscall/js"

	"interpreter/finder"
	"interpreter/session"
	"interpreter/shell"
	"interpreter/world"
)

var (
	w      world.World
	s      *session.Session
	vmSubs []js.Value
)

func main() {
	w = world.Load()
	s = session.New(w.Env)
	w.Finder.SetNotify(pushVM)

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
		w.Finder.SetViewport(args[0].Int(), args[1].Int())
		return nil
	})

	os["spawn"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		_, err := w.Finder.Spawn(args[0].String(), w.FS)
		if err != nil {
			return err.Error()
		}
		return nil
	})

	os["close"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.Finder.Close(args[0].String())
		return nil
	})

	os["closeFocused"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.Finder.CloseFocused()
		return nil
	})

	os["focus"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.Finder.Focus(args[0].String())
		return nil
	})

	os["drag"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.Finder.Drag(args[0].String(), args[1].Int(), args[2].Int())
		return nil
	})

	os["navigate"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		err := w.Finder.Navigate(args[0].String(), args[1].String(), w.FS)
		if err != nil {
			return err.Error()
		}
		return nil
	})

	os["back"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.Finder.Back(args[0].String())
		return nil
	})

	os["forward"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.Finder.Forward(args[0].String())
		return nil
	})

	os["setViewMode"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.Finder.SetViewMode(args[0].String(), finder.ViewMode(args[1].String()))
		return nil
	})

	os["selectEntry"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.Finder.SelectEntry(args[0].String(), args[1].String())
		return nil
	})

	os["openEntry"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		err := w.Finder.OpenEntry(args[0].String(), args[1].String(), w.FS)
		if err != nil {
			return err.Error()
		}
		return nil
	})

	js.Global().Set("os", js.ValueOf(os))

	select {}
}

func pushVM() {
	vm := finder.BuildVM(w.Finder, w.FS)
	jsonStr := vm.JSON()
	for _, cb := range vmSubs {
		cb.Invoke(jsonStr)
	}
}
