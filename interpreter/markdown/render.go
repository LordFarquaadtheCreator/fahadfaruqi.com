// Package markdown renders a file's GFM into sanitised HTML, which is what every panel that shows a
// document needs: the Finder's content pane, the viewer, and Quick Look.
package markdown

import (
	"bytes"
	"regexp"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

var parser = goldmark.New(goldmark.WithExtensions(extension.GFM))
var sanitizer = bluemonday.UGCPolicy()

// Render converts GFM to sanitised HTML.
func Render(content string) string {
	var buf bytes.Buffer
	if err := parser.Convert([]byte(content), &buf); err != nil {
		return "<p>Error rendering markdown</p>"
	}
	return sanitizer.Sanitize(sanitizeImages(buf.String()))
}

// IsDocument reports whether a path is one this package renders. A .txt is shown too, as plain
// preformatted text rather than parsed.
func IsDocument(path string) bool {
	switch strings.ToLower(ext(path)) {
	case ".md", ".markdown":
		return true
	}
	return false
}

func ext(path string) string {
	if i := strings.LastIndex(path, "."); i >= 0 {
		return path[i:]
	}
	return ""
}

// sanitizeImages strips <img> tags whose src is not /static/, /images/, or http(s).
func sanitizeImages(html string) string {
	re := regexp.MustCompile(`<img[^>]*>`)
	return re.ReplaceAllStringFunc(html, func(tag string) string {
		m := regexp.MustCompile(`src=["']([^"']+)["']`).FindStringSubmatch(tag)
		if len(m) < 2 {
			return ""
		}
		src := m[1]
		if strings.HasPrefix(src, "/static/") || strings.HasPrefix(src, "/images/") ||
			strings.HasPrefix(src, "http://") || strings.HasPrefix(src, "https://") {
			return tag
		}
		return ""
	})
}
