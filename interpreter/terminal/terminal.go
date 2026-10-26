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
	CWD    string `json:"cwd"`
}

// TerminalState holds per-window terminal state.
type TerminalState struct {
	Session *session.Session
	History []TerminalEntry
	// OpenURL is set by a script that opens a page: the window that ran it asks the browser for a
	// tab. It is the script's effect, so it belongs to the state, not to the view.
	OpenURL string
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

// SpawnScript opens a terminal window that has already run an executable shell script: one entry
// of history in the window's own voice, and the page the script asks for.
func (tm *TerminalManager) SpawnScript(scriptPath, url string) *wm.Window {
	win := tm.Spawn()
	state := tm.states[win.ID]
	state.OpenURL = url
	state.History = append(state.History, TerminalEntry{
		Input:  scriptPath,
		Output: textResponse("Opened " + url),
		CWD:    "~",
	})
	return win
}

// textResponse is the shell's own "open" response shape, so the transcript renders like any other
// command's output.
func textResponse(text string) string {
	b, _ := json.Marshal(struct {
		Type    string            `json:"type"`
		Payload map[string]string `json:"payload"`
		CWD     string            `json:"cwd"`
	}{Type: "open", Payload: map[string]string{"text": text}, CWD: "~"})
	return string(b)
}

func (tm *TerminalManager) RunCommand(id, input string) string {
	s, ok := tm.states[id]
	if !ok {
		return `{"type":"error","payload":{"command":"","message":"terminal not found"}}`
	}
	cwd := s.Session.CWD
	output := tm.runner.RunCommand(input, s.Session)

	// Check if command was clear — wipe history instead of recording
	var resp struct {
		Type string `json:"type"`
	}
	if json.Unmarshal([]byte(output), &resp) == nil && resp.Type == "clear" {
		s.History = make([]TerminalEntry, 0)
		return output
	}

	s.History = append(s.History, TerminalEntry{Input: input, Output: output, CWD: cwd})
	return output
}

// Interrupt records the line a visitor abandoned with Ctrl-C: the prompt as they had typed it, then
// the shell's own ^C, which is how a shell that never ran the command would read.
func (tm *TerminalManager) Interrupt(id, input string) {
	s, ok := tm.states[id]
	if !ok {
		return
	}
	s.History = append(s.History, TerminalEntry{Input: input, Output: echoText("^C"), CWD: s.Session.CWD})
}

func echoText(text string) string {
	b, _ := json.Marshal(struct {
		Type    string            `json:"type"`
		Payload map[string]string `json:"payload"`
	}{Type: "echo", Payload: map[string]string{"text": text}})
	return string(b)
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
	OpenURL string          `json:"openUrl,omitempty"`
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
		OpenURL: s.OpenURL,
	}
	b, _ := json.Marshal(vm)
	return b
}
