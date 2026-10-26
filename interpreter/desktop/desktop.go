package desktop

import (
	"encoding/json"
	"interpreter/apps"
	"interpreter/assets"
	"interpreter/fs"
	"path"
)

// DesktopVM is the desktop icons view model.
type DesktopVM struct {
	Icons     []DesktopIconVM `json:"icons"`
	Wallpaper string          `json:"wallpaper"`
}

type DesktopIconVM struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Icon string `json:"icon"` // icon type: "folder", "file", "image", "text"
	X    int    `json:"x"`
	Y    int    `json:"y"`
}

// BuildVM builds the desktop view model from the filesystem.
// Shows top-level non-hidden items in home directory as desktop icons.
// dark selects the wallpaper variant.
func BuildVM(filesystem fs.FileSystem, dark bool) DesktopVM {
	children, paths, _ := filesystem.Children("~", false)
	var icons []DesktopIconVM

	col := 0
	row := 0
	for i, child := range children {
		iconPath := assets.URL("/icons/file.png")
		if child.Type.IsDir() {
			iconPath = assets.URL("/icons/folder.png")
		} else if child.Meta.MimeType != "" {
			if isImageType(child.Meta.MimeType) {
				iconPath = assets.URL("/icons/file.png")
			}
		}

		icons = append(icons, DesktopIconVM{
			Name: path.Base(paths[i]),
			Path: paths[i],
			Icon: iconPath,
			X:    20 + col*100, // right-aligned via CSS right offset
			Y:    40 + row*100,
		})

		row++
		if row > 5 {
			row = 0
			col++
		}
	}

	// Apps that open from the desktop rather than the dock are appended after the filesystem's own
	// entries, which puts them below the last folder.
	if app, ok := apps.Get("rose"); ok {
		icons = append(icons, DesktopIconVM{
			Name: app.Name,
			Path: apps.AppPathPrefix + app.ID,
			Icon: app.Icon,
			X:    20 + col*100,
			Y:    40 + row*100,
		})
	}

	wp := assets.URL("/wallpapers/compressed/wallpaper.avif")
	if dark {
		wp = assets.URL("/wallpapers/compressed/wallpaper-dark.avif")
	}

	return DesktopVM{Icons: icons, Wallpaper: wp}
}

func isImageType(mime string) bool {
	return mime == "image/png" || mime == "image/jpeg" || mime == "image/gif" || mime == "image/webp"
}

// JSON marshals DesktopVM to JSON.
func (vm DesktopVM) JSON() string {
	b, err := json.Marshal(vm)
	if err != nil {
		return `{"icons":[]}`
	}
	return string(b)
}
