package markdown

import (
	"strings"
	"testing"
)

func TestRenderMakesAPage(t *testing.T) {
	html := Render("# About\n\n**Name:** Fahad\n\n- one\n- two")

	for _, want := range []string{"<h1>About</h1>", "<strong>Name:</strong>", "<li>one</li>"} {
		if !strings.Contains(html, want) {
			t.Errorf("render missed %q in:\n%s", want, html)
		}
	}
}

// Everything the visitor reads is content, so a document cannot bring a script with it.
func TestRenderStripsScripts(t *testing.T) {
	html := Render("# hi\n\n<script>alert(1)</script>\n\n[bad](javascript:alert(1))")

	if strings.Contains(strings.ToLower(html), "<script") {
		t.Errorf("script survived sanitising:\n%s", html)
	}
	if strings.Contains(html, "javascript:") {
		t.Errorf("javascript: URL survived sanitising:\n%s", html)
	}
}

func TestRenderKeepsAllowedImages(t *testing.T) {
	html := Render("![ok](https://example.com/a.png) ![local](/images/b.png) ![no](file:///etc/passwd)")

	if !strings.Contains(html, "https://example.com/a.png") || !strings.Contains(html, "/images/b.png") {
		t.Errorf("allowed images were dropped:\n%s", html)
	}
	if strings.Contains(html, "file:///") {
		t.Errorf("a local file image survived:\n%s", html)
	}
}

func TestIsDocument(t *testing.T) {
	for _, path := range []string{"~/a.md", "~/B.MARKDOWN", "x.md"} {
		if !IsDocument(path) {
			t.Errorf("IsDocument(%q) = false", path)
		}
	}
	for _, path := range []string{"~/gallery.sh", "~/notes.txt", "~"} {
		if IsDocument(path) {
			t.Errorf("IsDocument(%q) = true", path)
		}
	}
}
