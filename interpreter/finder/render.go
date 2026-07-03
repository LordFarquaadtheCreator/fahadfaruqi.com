package finder

import (
	"bytes"
	"regexp"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

var mdParser = goldmark.New(goldmark.WithExtensions(extension.GFM))
var sanitizer = bluemonday.UGCPolicy()

// RenderMarkdown converts GFM to sanitized HTML.
func RenderMarkdown(content string) string {
	var buf bytes.Buffer
	if err := mdParser.Convert([]byte(content), &buf); err != nil {
		return "<p>Error rendering markdown</p>"
	}
	html := sanitizeImages(buf.String())
	return sanitizer.Sanitize(html)
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
