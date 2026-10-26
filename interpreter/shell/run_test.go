package shell

import (
	"encoding/json"
	"strings"
	"testing"

	"interpreter/fs"
	"interpreter/session"
	"interpreter/terminal"
	"interpreter/world"
)

// The shell is exercised through a terminal in the real world — the same filesystem and the same
// commands a visitor gets — so a passing test is a passing keystroke, not a mocked one.
// the runner holds the world by pointer: the terminal is installed into it after the manager is
// built, and the shell has to see that terminal when a command asks for a page
type runner struct{ w *world.World }

func (r runner) RunCommand(input string, s *session.Session) string { return Run(input, *r.w, s) }
func (r runner) TabComplete(input, cwd string) []string             { return nil }

type reply struct {
	Type    string `json:"type"`
	Payload struct {
		Text    string `json:"text"`
		Message string `json:"message"`
		Command string `json:"command"`
	} `json:"payload"`
	CWD string `json:"cwd"`
}

func machine(t *testing.T) (world.World, *terminal.TerminalManager, string) {
	t.Helper()
	w := world.Load()
	tm := terminal.NewTerminalManager(w.WM, runner{&w}, w.Env)
	w.Terminal = tm
	win := tm.Spawn()
	return w, tm, win.ID
}

// type runs a line in a terminal of the real world and decodes what the window would show.
func typeInto(t *testing.T, tm *terminal.TerminalManager, id, line string) reply {
	t.Helper()
	var r reply
	if err := json.Unmarshal([]byte(tm.RunCommand(id, line)), &r); err != nil {
		t.Fatalf("run %q: %v", line, err)
	}
	return r
}

func TestScriptRunsFromItsPath(t *testing.T) {
	_, tm, id := machine(t)

	got := typeInto(t, tm, id, "./Desktop/gallery.sh")
	if got.Type != "open" {
		t.Fatalf("./Desktop/gallery.sh: type %q, want open (payload %+v)", got.Type, got.Payload)
	}
	if !strings.Contains(got.Payload.Text, "https://art.fahadfaruqi.com") {
		t.Errorf("./Desktop/gallery.sh opened %q, want the gallery", got.Payload.Text)
	}
	if state, ok := tm.GetState(id); !ok || state.OpenURL != "https://art.fahadfaruqi.com" {
		t.Errorf("window was not asked to open the page: %+v", state)
	}
}

func TestScriptRunsRelativeToTheShellsOwnDirectory(t *testing.T) {
	_, tm, id := machine(t)

	if got := typeInto(t, tm, id, "cd Desktop"); got.Type != "noop" && got.Type != "varset" {
		t.Fatalf("cd Desktop: %q", got.Type)
	}
	got := typeInto(t, tm, id, "./gallery.sh")
	if got.Type != "open" || !strings.Contains(got.Payload.Text, "art.fahadfaruqi.com") {
		t.Fatalf("./gallery.sh inside Desktop: %+v", got)
	}
}

func TestBareScriptNameIsNotACommand(t *testing.T) {
	_, tm, id := machine(t)

	got := typeInto(t, tm, id, "gallery.sh")
	if got.Type != "error" || !strings.Contains(got.Payload.Message, "command not found") {
		t.Fatalf("gallery.sh: %+v, want command not found", got)
	}
}

func TestScriptPathMustExist(t *testing.T) {
	_, tm, id := machine(t)

	got := typeInto(t, tm, id, "./Desktop/missing.sh")
	if got.Type != "error" || !strings.Contains(got.Payload.Message, "no such file or directory") {
		t.Fatalf("missing script: %+v", got)
	}
}

func TestAPlainFileIsNotExecutable(t *testing.T) {
	_, tm, id := machine(t)

	got := typeInto(t, tm, id, "./Desktop/About/about.md")
	if got.Type != "error" || !strings.Contains(got.Payload.Message, "permission denied") {
		t.Fatalf("markdown executed: %+v", got)
	}
}

func TestADirectoryIsNotExecutable(t *testing.T) {
	_, tm, id := machine(t)

	got := typeInto(t, tm, id, "./Desktop")
	if got.Type != "error" || !strings.Contains(got.Payload.Message, "is a directory") {
		t.Fatalf("directory executed: %+v", got)
	}
}

func TestSourceRunsAScriptInTheWindow(t *testing.T) {
	for _, line := range []string{
		"source Desktop/gallery.sh",
		"source ./Desktop/gallery.sh",
		". Desktop/gallery.sh",
	} {
		t.Run(line, func(t *testing.T) {
			_, tm, id := machine(t)
			got := typeInto(t, tm, id, line)
			if got.Type != "open" {
				t.Fatalf("%s: %+v", line, got)
			}
			if state, _ := tm.GetState(id); state.OpenURL != "https://art.fahadfaruqi.com" {
				t.Errorf("%s: page not opened (%+v)", line, state)
			}
		})
	}
}

func TestSourceNeedsAFile(t *testing.T) {
	_, tm, id := machine(t)

	if got := typeInto(t, tm, id, "source nope.sh"); got.Type != "error" ||
		!strings.Contains(got.Payload.Message, "no such file or directory") {
		t.Fatalf("source nope.sh: %+v", got)
	}
	if got := typeInto(t, tm, id, "source"); got.Type != "error" ||
		!strings.Contains(got.Payload.Message, "usage") {
		t.Fatalf("source with no argument: %+v", got)
	}
}

func TestOpenURLAsksTheWindowForThePage(t *testing.T) {
	_, tm, id := machine(t)

	got := typeInto(t, tm, id, "open https://example.com/x")
	if got.Type != "open" {
		t.Fatalf("open url: %+v", got)
	}
	if state, _ := tm.GetState(id); state.OpenURL != "https://example.com/x" {
		t.Errorf("open url recorded %q", state.OpenURL)
	}
}

func TestScriptLinesSkipWhatsNotACommand(t *testing.T) {
	got := scriptLines("#!/bin/zsh\r\n# a comment\r\n\r\nopen https://example.com\r\n  \r\n")
	if len(got) != 1 || got[0] != "open https://example.com" {
		t.Fatalf("scriptLines = %q", got)
	}
}

// A script that sources itself has to stop somewhere; the shell reports the depth rather than
// running out of stack.
func TestAScriptCannotRunItselfForever(t *testing.T) {
	w := world.World{
		FS:  fs.NewFileSystem(loopNodes()),
		Env: map[string]string{"HOME": "~"},
	}
	s := session.New(w.Env)

	var r reply
	if err := json.Unmarshal([]byte(Run("source ~/loop.sh", w, s)), &r); err != nil {
		t.Fatal(err)
	}
	if r.Type != "error" || !strings.Contains(r.Payload.Message, "maximum nesting depth") {
		t.Fatalf("self-sourcing script: %+v", r)
	}
}

func loopNodes() map[string]fs.Node {
	return map[string]fs.Node{
		"~":         {Type: fs.Dir, Children: []string{"~/loop.sh"}},
		"~/loop.sh": {Type: fs.File, Content: "source ~/loop.sh\n"},
	}
}

// A sourced file keeps its variables: that is the difference between sourcing and running it.
func TestSourceKeepsTheVariablesItSets(t *testing.T) {
	nodes := loopNodes()
	nodes["~/env.sh"] = fs.Node{Type: fs.File, Content: "# sets a variable\nfoo=bar\n"}
	home := nodes["~"]
	home.Children = append(home.Children, "~/env.sh")
	nodes["~"] = home
	w := world.World{FS: fs.NewFileSystem(nodes), Env: map[string]string{"HOME": "~"}}
	s := session.New(w.Env)

	Run("source ~/env.sh", w, s)
	if s.Vars["foo"] != "bar" {
		t.Fatalf("variables after source: %+v", s.Vars)
	}
}
