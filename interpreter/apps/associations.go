package apps

import (
	"path"
	"strings"
)

// Extensions maps file extensions to app IDs.
var Extensions = map[string]string{
	".png":  "viewer-image",
	".jpeg": "viewer-image",
	".jpg":  "viewer-image",
	".gif":  "viewer-image",
	".webp": "viewer-image",
	".md":   "viewer-text",
	".txt":  "viewer-text",
}

// AppForFile resolves which app should open a given filename based on its extension.
func AppForFile(filename string) string {
	ext := strings.ToLower(path.Ext(filename))
	if appID, ok := Extensions[ext]; ok {
		return appID
	}
	return ""
}
