package ephemeral_test

import (
	"encoding/json"
	"strings"
	"testing"

	"interpreter/ephemeral"
	"interpreter/world"
)

// Quick Look is exercised against the real filesystem, so a passing test is what a visitor sees when
// they press space on a file.
func machine(t *testing.T) (world.World, *ephemeral.EphemeralManager) {
	t.Helper()
	w := world.Load()
	return w, ephemeral.NewEphemeralManager(w.WM, w.FS)
}

func look(t *testing.T, w world.World, em *ephemeral.EphemeralManager, path string) ephemeral.EphemeralContentVM {
	t.Helper()
	win := em.SpawnForFile(path)
	raw := ephemeral.BuildContentVM(win, em)

	var vm ephemeral.EphemeralContentVM
	if err := json.Unmarshal(raw, &vm); err != nil {
		t.Fatalf("decode %s: %v", path, err)
	}
	return vm
}

func TestQuickLookRendersMarkdown(t *testing.T) {
	w, em := machine(t)

	vm := look(t, w, em, "~/Desktop/About/about.md")
	if vm.Kind != "ephemeral-markdown" {
		t.Fatalf("about.md previewed as %q, want ephemeral-markdown", vm.Kind)
	}
	if !strings.Contains(vm.HTML, "<h1>About</h1>") {
		t.Errorf("the heading was not rendered: %q", vm.HTML)
	}
	if !strings.Contains(vm.HTML, "Fahad Faruqi") {
		t.Errorf("the body was not rendered: %q", vm.HTML)
	}
	if vm.Name != "about.md" || vm.FilePath != "~/Desktop/About/about.md" {
		t.Errorf("preview names %q at %q", vm.Name, vm.FilePath)
	}
}

func TestQuickLookFitsADocumentSheet(t *testing.T) {
	w, em := machine(t)

	before := len(w.WM.Windows())
	win := em.SpawnForFile("~/Desktop/Resume/experience/tobi-wealth.md")
	if got := len(w.WM.Windows()); got != before+1 {
		t.Fatalf("Quick Look did not open a window: %d → %d", before, got)
	}
	if win.Width <= 0 || win.Height <= 0 {
		t.Fatalf("opened at %dx%d", win.Width, win.Height)
	}
	// a page is taller than it is wide, and never wider than the sheet's 820pt cap
	if win.Width > 820 {
		t.Errorf("document sheet is %d wide, want at most 820", win.Width)
	}
	if win.Height <= win.Width {
		t.Errorf("document sheet is %dx%d, want a page", win.Width, win.Height)
	}
}

func TestQuickLookStillShowsFolders(t *testing.T) {
	w, em := machine(t)

	vm := look(t, w, em, "~/Desktop/About")
	if vm.Kind != "ephemeral-folder" || !strings.Contains(vm.Icon, "folder") {
		t.Errorf("folder previewed as %+v", vm)
	}
}

// A script has no picture and no page: Quick Look shows the file it is, not an empty frame.
func TestQuickLookShowsAPlainFile(t *testing.T) {
	w, em := machine(t)

	vm := look(t, w, em, "~/Desktop/gallery.sh")
	if vm.Kind != "ephemeral-folder" || !strings.Contains(vm.Icon, "file.png") {
		t.Errorf("script previewed as %+v", vm)
	}
	if vm.Src != "" {
		t.Errorf("script preview carries an image source: %q", vm.Src)
	}
}

// Pressing space twice replaces the sheet rather than stacking previews.
func TestQuickLookReplacesTheOpenSheet(t *testing.T) {
	w, em := machine(t)

	first := em.SpawnForFile("~/Desktop/About/about.md")
	em.SpawnForFile("~/Desktop/Contact/contact.md")
	for _, win := range w.WM.Windows() {
		if win.ID == first.ID {
			t.Fatalf("the first sheet was left open")
		}
	}
}
