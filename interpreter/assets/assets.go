// Package assets resolves asset paths to the origin they are served from.
//
// The site's images, icons and wallpapers are not in the repo: they live in the
// Cloudflare R2 bucket served at https://assets.fahadfaruqi.com. Base is baked
// into the wasm at build time and can be overridden with
//
//	go build -ldflags "-X interpreter/assets.Base=<origin>"
package assets

import "strings"

// Base is the origin every asset path is prefixed with.
var Base = "https://assets.fahadfaruqi.com/www"

// URL returns the absolute URL for p, an absolute asset path such as
// "/icons/finder.png". An empty or already-absolute path passes through
// unchanged, so wrapping a value twice is harmless.
func URL(p string) string {
	if p == "" || strings.HasPrefix(p, "http://") || strings.HasPrefix(p, "https://") {
		return p
	}
	return Base + p
}
