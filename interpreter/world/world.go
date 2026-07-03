package world

import (
	"interpreter/finder"
	"interpreter/fs"
)

type World struct {
	FS     fs.FileSystem
	Env    map[string]string
	Finder *finder.Manager
}
