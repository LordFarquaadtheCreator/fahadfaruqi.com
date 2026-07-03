package terminal

import (
	"encoding/json"
	"interpreter/apps"
	"interpreter/session"
	"interpreter/wm"
)

// CommandRunner runs shell commands. Implemented by world.World to avoid import cycles.
type CommandRunner interface {
	RunCommand(input string, s *session.Session) string
	TabComplete(input, cwd string) []string
}

// TerminalEntry is one command + output pair in terminal history.
type TerminalEntry struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

// TerminalState holds per-window terminal state.
type TerminalState struct {
	Session *session.Session
	History []TerminalEntry
}

// TerminalManager owns per-window terminal state.
type TerminalManager struct {
	wm     *wm.Manager
	states map[string]*TerminalState
	runner CommandRunner
	env    map[string]string
}

func NewTerminalManager(w *wm.Manager, runner CommandRunner, env map[string]string) *TerminalManager {
	return &TerminalManager{
		wm:     w,
		states: make(map[string]*TerminalState),
		runner: runner,
		env:    env,
	}
}

func (tm *TerminalManager) Spawn() *wm.Window {
	app, _ := apps.Get("terminal")
	win := tm.wm.SpawnSilent(string(apps.TypeTerminal), app.ID, app.Name, app.Width, app.Height)
	tm.states[win.ID] = &TerminalState{
		Session: session.New(tm.env),
		History: make([]TerminalEntry, 0),
	}
	return win
}

func (tm *TerminalManager) RunCommand(id, input string) string {
	s, ok := tm.states[id]
	if !ok {
		return `{"type":"error","payload":{"command":"","message":"terminal not found"}}`
	}
	output := tm.runner.RunCommand(input, s.Session)
	s.History = append(s.History, TerminalEntry{Input: input, Output: output})
	return output
}

func (tm *TerminalManager) TabComplete(id, input string) []string {
	s, ok := tm.states[id]
	if !ok {
		return nil
	}
	return tm.runner.TabComplete(input, s.Session.CWD)
}

func (tm *TerminalManager) GetState(id string) (*TerminalState, bool) {
	s, ok := tm.states[id]
	return s, ok
}

func (tm *TerminalManager) Cleanup(id string) {
	delete(tm.states, id)
}

// TerminalContentVM is the app-specific content for a Terminal window.
type TerminalContentVM struct {
	Kind    string          `json:"kind"` // "terminal"
	CWD     string          `json:"cwd"`
	History []TerminalEntry `json:"history"`
}

// BuildContentVM builds terminal-specific content for a window.
func BuildContentVM(win *wm.Window, tm *TerminalManager) json.RawMessage {
	s, ok := tm.GetState(win.ID)
	if !ok {
		return json.RawMessage(`{"kind":"terminal","cwd":"~","history":[]}`)
	}
	vm := TerminalContentVM{
		Kind:    "terminal",
		CWD:     s.Session.CWD,
		History: s.History,
	}
	b, _ := json.Marshal(vm)
	return b
}
