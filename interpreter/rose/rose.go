// Package rose owns the Rose app's state: which drawing a window is showing.
//
// The frontend draws the rose with a canvas, but it does not decide what the rose is — the seed
// lives here, so a shuffle is a VM change like any other and the window can be rebuilt from the
// JSON alone.
package rose

import (
	"encoding/json"
	"math/rand"

	"interpreter/apps"
	"interpreter/wm"
)

// baseSeed is the sketch's own seed: the first window always opens on the drawing as published.
const baseSeed = 7

// RoseState is per-window state.
type RoseState struct {
	Seed int
}

// RoseManager owns per-window rose state.
type RoseManager struct {
	wm     *wm.Manager
	states map[string]*RoseState
}

func NewRoseManager(w *wm.Manager) *RoseManager {
	return &RoseManager{
		wm:     w,
		states: make(map[string]*RoseState),
	}
}

// Spawn opens a rose window on the base seed.
func (rm *RoseManager) Spawn() *wm.Window {
	app, _ := apps.Get("rose")
	win := rm.wm.SpawnSilent(string(app.Type), app.ID, app.Name, app.Width, app.Height)
	rm.states[win.ID] = &RoseState{Seed: baseSeed}
	return win
}

// Shuffle gives one rose window a new seed. It reports whether the window was a rose window,
// so the caller only pushes a VM when something actually changed.
func (rm *RoseManager) Shuffle(id string) bool {
	state, ok := rm.states[id]
	if !ok {
		return false
	}
	seed := baseSeed
	for seed == state.Seed {
		seed = rand.Intn(1 << 20)
	}
	state.Seed = seed
	return true
}

func (rm *RoseManager) Cleanup(id string) {
	delete(rm.states, id)
}

// ContentVM is the rose window's content.
type ContentVM struct {
	Kind string `json:"kind"` // "rose"
	Seed int    `json:"seed"`
}

func BuildContentVM(win *wm.Window, rm *RoseManager) json.RawMessage {
	state, ok := rm.states[win.ID]
	if !ok {
		state = &RoseState{Seed: baseSeed}
		rm.states[win.ID] = state
	}
	b, err := json.Marshal(ContentVM{Kind: "rose", Seed: state.Seed})
	if err != nil {
		return json.RawMessage(`{"kind":"rose","seed":7}`)
	}
	return b
}
