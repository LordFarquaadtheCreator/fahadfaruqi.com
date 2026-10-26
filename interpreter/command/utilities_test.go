package command

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
)

// The utilities are exercised against the machine the visitor gets: the real filesystem, the real
// environment, read through a session standing in the home folder.
func machine(t *testing.T) (world.World, *session.Session) {
	t.Helper()
	w := world.Load()
	s := session.New(w.Env)
	return w, s
}

// run calls a command the way the shell does and returns the text it printed.
func run(t *testing.T, w world.World, s *session.Session, cmd string, args ...string) string {
	t.Helper()
	fn, ok := Registry[cmd]
	if !ok {
		t.Fatalf("%s is not registered", cmd)
	}
	out := response.Marshal(fn(args, w, s))
	var r struct {
		Type    string `json:"type"`
		Payload struct {
			Text    string `json:"text"`
			Message string `json:"message"`
		} `json:"payload"`
	}
	if err := json.Unmarshal([]byte(out), &r); err != nil {
		t.Fatalf("%s: %v", cmd, err)
	}
	if r.Type == string(response.Err) {
		return "error: " + r.Payload.Message
	}
	return r.Payload.Text
}

// The registry and the list which uses have to agree, or a command works but denies it exists.
func TestBuiltinsMatchTheRegistry(t *testing.T) {
	inList := make(map[string]bool, len(builtins))
	for _, name := range builtins {
		if inList[name] {
			t.Errorf("builtins lists %q twice", name)
		}
		inList[name] = true
	}
	for name := range Registry {
		if !inList[name] {
			t.Errorf("registry has %q, builtins does not", name)
		}
	}
	for name := range inList {
		if _, ok := Registry[name]; !ok {
			t.Errorf("builtins has %q, registry does not", name)
		}
	}
}

func TestHeadTakesTheFirstLines(t *testing.T) {
	w, s := machine(t)

	if got := strings.Count(run(t, w, s, "head", "~/Desktop/gallery.sh"), "\n"); got != 2 {
		t.Errorf("head of a three-line file returned %d newlines: %q", got, got)
	}
	if got := run(t, w, s, "head", "-n", "1", "~/Desktop/gallery.sh"); got != "#!/bin/zsh" {
		t.Errorf("head -n 1 = %q", got)
	}
	if got := run(t, w, s, "head", "-n1", "~/Desktop/gallery.sh"); got != "#!/bin/zsh" {
		t.Errorf("head -n1 = %q", got)
	}
	if got := run(t, w, s, "head", "--lines=1", "~/Desktop/gallery.sh"); got != "#!/bin/zsh" {
		t.Errorf("head --lines=1 = %q", got)
	}
	if got := run(t, w, s, "head", "-n", "x", "~/Desktop/gallery.sh"); !strings.HasPrefix(got, "error: invalid number") {
		t.Errorf("head -n x = %q", got)
	}
	if got := run(t, w, s, "head"); !strings.HasPrefix(got, "error: usage") {
		t.Errorf("head with no file = %q", got)
	}
	if got := run(t, w, s, "head", "~/Desktop/About"); !strings.Contains(got, "is a directory") {
		t.Errorf("head of a folder = %q", got)
	}
}

func TestTailTakesTheLastLines(t *testing.T) {
	w, s := machine(t)

	got := run(t, w, s, "tail", "-n", "1", "~/Desktop/gallery.sh")
	if got != "open https://art.fahadfaruqi.com" {
		t.Errorf("tail -n 1 = %q", got)
	}
	// fewer lines asked for than the file has: the whole file comes back
	if got := run(t, w, s, "tail", "-n", "99", "~/Desktop/gallery.sh"); !strings.Contains(got, "#!/bin/zsh") {
		t.Errorf("tail -n 99 = %q", got)
	}
}

func TestWCCountsWhatTheFileHolds(t *testing.T) {
	w, s := machine(t)

	node, err := w.FS.Stat("~/Desktop/gallery.sh")
	if err != nil {
		t.Fatal(err)
	}
	want := fmt.Sprintf("%d %d %d ~/Desktop/gallery.sh",
		len(contentLines(node.Content)), len(strings.Fields(node.Content)), len(node.Content))
	if got := run(t, w, s, "wc", "~/Desktop/gallery.sh"); got != want {
		t.Errorf("wc = %q, want %q", got, want)
	}
	if got := run(t, w, s, "wc", "-l", "~/Desktop/gallery.sh"); got != "3 ~/Desktop/gallery.sh" {
		t.Errorf("wc -l = %q", got)
	}
	if got := run(t, w, s, "wc", "-c", "~/Desktop/gallery.sh"); got != "82 ~/Desktop/gallery.sh" {
		t.Errorf("wc -c = %q", got)
	}
}

func TestGrepSearchesLines(t *testing.T) {
	w, s := machine(t)
	file := "~/Desktop/Resume/toolkit/languages.md"

	if got := run(t, w, s, "grep", "Swift", file); got != "- Swift" {
		t.Errorf("grep Swift = %q", got)
	}
	if got := run(t, w, s, "grep", "swift", file); got != "" {
		t.Errorf("grep is case-sensitive by default, got %q", got)
	}
	if got := run(t, w, s, "grep", "-i", "swift", file); got != "- Swift" {
		t.Errorf("grep -i = %q", got)
	}
	if got := run(t, w, s, "grep", "-n", "^# ", file); got != "1:# Languages" {
		t.Errorf("grep -n = %q", got)
	}
	if got := run(t, w, s, "grep", "-c", "^- ", file); !regexp.MustCompile(`^\d+$`).MatchString(got) {
		t.Errorf("grep -c = %q", got)
	}
	if got := run(t, w, s, "grep", "-v", "^#", "~/Desktop/Contact/contact.md"); strings.Contains(got, "# Contact") {
		t.Errorf("grep -v kept the header: %q", got)
	}
	if got := run(t, w, s, "grep", "[", file); !strings.HasPrefix(got, "error: invalid pattern") {
		t.Errorf("grep with a broken pattern = %q", got)
	}
	if got := run(t, w, s, "grep", "nothing-matches-this", file); got != "" {
		t.Errorf("grep with no match printed %q, want silence", got)
	}
}

func TestFindWalksTheTree(t *testing.T) {
	w, s := machine(t)

	got := run(t, w, s, "find", "~/Desktop/Resume", "-name", "*.md")
	for _, want := range []string{"~/Desktop/Resume/toolkit/languages.md", "~/Desktop/Resume/experience/google.md"} {
		if !strings.Contains(got, want) {
			t.Errorf("find missed %s in:\n%s", want, got)
		}
	}
	if strings.Contains(got, "gallery.sh") {
		t.Errorf("find -name '*.md' matched a script:\n%s", got)
	}
	if got := run(t, w, s, "find", "~/Desktop/gallery.sh", "-name", "*.sh"); got != "~/Desktop/gallery.sh" {
		t.Errorf("find on one file = %q", got)
	}
}

func TestTreeShowsTheFolders(t *testing.T) {
	w, s := machine(t)

	got := run(t, w, s, "tree", "~/Desktop")
	if !strings.HasPrefix(got, "~/Desktop") {
		t.Errorf("tree did not start at its path: %q", got)
	}
	for _, want := range []string{"├── About/", "├── Resume/", "└── gallery.sh", "experience/"} {
		if !strings.Contains(got, want) {
			t.Errorf("tree missed %q in:\n%s", want, got)
		}
	}
	if !strings.Contains(got, "gallery.sh") {
		t.Errorf("tree missed the script:\n%s", got)
	}
}

func TestStatReportsAPath(t *testing.T) {
	w, s := machine(t)

	got := run(t, w, s, "stat", "~/Desktop/gallery.sh")
	for _, want := range []string{"path: ~/Desktop/gallery.sh", "type: file", "size: 82"} {
		if !strings.Contains(got, want) {
			t.Errorf("stat missed %q in:\n%s", want, got)
		}
	}
	if got := run(t, w, s, "stat", "~/Desktop"); !strings.Contains(got, "type: directory") {
		t.Errorf("stat of a folder = %q", got)
	}
	if got := run(t, w, s, "stat", "~/nope"); !strings.Contains(got, "no such file") {
		t.Errorf("stat of a missing path = %q", got)
	}
}

func TestWhichFindsBuiltins(t *testing.T) {
	w, s := machine(t)

	if got := run(t, w, s, "which", "grep"); got != "/usr/bin/grep" {
		t.Errorf("which grep = %q", got)
	}
	if got := run(t, w, s, "which", "grep", "nope"); !strings.Contains(got, "nope not found") {
		t.Errorf("which nope = %q", got)
	}
}

// A Mac in New York names its zone; the wasm runtime only reports an offset, so the command has to
// translate it rather than print "UTC-4" in a menu bar that says EDT.
func TestDateNamesTheZone(t *testing.T) {
	if got := zoneName(time.Date(2026, 10, 26, 12, 0, 0, 0, time.UTC)); got != "UTC" {
		t.Errorf("zoneName(UTC) = %q", got)
	}
}

func TestDateHostnameAndUname(t *testing.T) {
	w, s := machine(t)

	if got := run(t, w, s, "date"); !regexp.MustCompile(`^\w{3} \w{3} +\d+ \d\d:\d\d:\d\d \w+ \d{4}$`).MatchString(got) {
		t.Errorf("date = %q", got)
	}
	if got := run(t, w, s, "whoami"); got != "fahad" {
		t.Errorf("whoami = %q, want the prompt's user", got)
	}
	if got := run(t, w, s, "hostname"); got != "mac" {
		t.Errorf("hostname = %q, want the prompt's host", got)
	}
	if got := run(t, w, s, "uname"); got != "Darwin" {
		t.Errorf("uname = %q", got)
	}
	if got := run(t, w, s, "uname", "-a"); !strings.HasPrefix(got, "Darwin mac 24.0.0 arm64") {
		t.Errorf("uname -a = %q, want the machine's own name and arch", got)
	}
}
