package desktop

import (
	"encoding/json"
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
		iconPath := "/icons/file.png"
		if child.Type.IsDir() {
			iconPath = "/icons/folder.png"
		} else if child.Meta.MimeType != "" {
			if isImageType(child.Meta.MimeType) {
				iconPath = "/icons/file.png"
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

	wp := "/wallpapers/wallpaper.jpg"
	if dark {
		wp = "/wallpapers/wallpaper-dark.jpg"
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
