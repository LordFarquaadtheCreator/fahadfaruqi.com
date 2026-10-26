package terminal_test

import (
	"slices"
	"testing"

	"interpreter/shell"
	"interpreter/terminal"
	"interpreter/world"
)

// Completion offers what the shell can run, so a new command is completable the moment it is
// registered — there is no second list to keep in step.
func TestCompletionFollowsTheShellsCommands(t *testing.T) {
	w := world.Load()

	for _, want := range []string{"grep", "tree", "stat", "which", "uname"} {
		if got := terminal.Complete(want[:2], "~", w.FS, shell.Commands()); !slices.Contains(got, want) {
			t.Errorf("completing %q did not offer %q: %v", want[:2], want, got)
		}
	}
	if got := terminal.Complete("", "~", w.FS, shell.Commands()); got != nil {
		t.Errorf("empty input completed to %v", got)
	}
}
