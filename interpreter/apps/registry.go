package apps

import "interpreter/assets"

type AppType string

const (
	TypeFinder      AppType = "finder"
	TypeTerminal    AppType = "terminal"
	TypeViewer      AppType = "viewer"
	TypeEphemeral   AppType = "ephemeral"
	TypePlaceholder AppType = "placeholder"
	TypeEditor      AppType = "editor"
	TypeRose        AppType = "rose"
	TypeScript      AppType = "script"
)

type App struct {
	ID     string
	Name   string
	Icon   string
	Type   AppType
	Width  int
	Height int
}

var Registry = map[string]App{
	"finder":       {ID: "finder", Name: "Finder", Icon: assets.URL("/icons/finder.png"), Type: TypeFinder, Width: 720, Height: 480},
	"terminal":     {ID: "terminal", Name: "iTerm", Icon: assets.URL("/icons/iterm.png"), Type: TypeTerminal, Width: 640, Height: 400},
	"viewer-image": {ID: "viewer-image", Name: "Preview", Icon: assets.URL("/icons/preview.png"), Type: TypeViewer, Width: 600, Height: 450},
	"viewer-text":  {ID: "viewer-text", Name: "Preview", Icon: assets.URL("/icons/preview.png"), Type: TypeViewer, Width: 600, Height: 450},
	"preview-ephemeral": {ID: "preview-ephemeral", Name: "Quick Look", Icon: assets.URL("/icons/preview.png"), Type: TypeEphemeral, Width: 400, Height: 300},
	"obsidian":     {ID: "obsidian", Name: "Obsidian", Icon: assets.URL("/icons/obsidian.png"), Type: TypePlaceholder, Width: 800, Height: 500},
	"slack":        {ID: "slack", Name: "Slack", Icon: assets.URL("/icons/slack.png"), Type: TypePlaceholder, Width: 800, Height: 500},
	"spotify":      {ID: "spotify", Name: "Spotify", Icon: assets.URL("/icons/spotify.png"), Type: TypePlaceholder, Width: 800, Height: 500},
	"steam":        {ID: "steam", Name: "Steam", Icon: assets.URL("/icons/steam.png"), Type: TypePlaceholder, Width: 800, Height: 500},
	"devin":        {ID: "devin", Name: "Devin", Icon: assets.URL("/icons/devin.png"), Type: TypeEditor, Width: 900, Height: 600},
	"zed":          {ID: "zed", Name: "Zed", Icon: assets.URL("/icons/zed.png"), Type: TypeEditor, Width: 900, Height: 600},
	"xcode":        {ID: "xcode", Name: "Xcode", Icon: assets.URL("/icons/xcode.png"), Type: TypeEditor, Width: 900, Height: 600},
	"figma":        {ID: "figma", Name: "Figma", Icon: assets.URL("/icons/figma.png"), Type: TypeEditor, Width: 900, Height: 600},
	"hermes":       {ID: "hermes", Name: "Hermes", Icon: assets.URL("/icons/hermes.png"), Type: TypeEditor, Width: 900, Height: 600},
	"affinity":     {ID: "affinity", Name: "Affinity Designer", Icon: assets.URL("/icons/affinity.png"), Type: TypePlaceholder, Width: 900, Height: 600},
	"whatsapp":     {ID: "whatsapp", Name: "WhatsApp", Icon: assets.URL("/icons/whatsapp.png"), Type: TypePlaceholder, Width: 800, Height: 500},
	"imsg":         {ID: "imsg", Name: "Messages", Icon: assets.URL("/icons/imsg.png"), Type: TypePlaceholder, Width: 800, Height: 500},
	"browseros":    {ID: "browseros", Name: "BrowserOS", Icon: assets.URL("/icons/browseros.png"), Type: TypePlaceholder, Width: 1000, Height: 650},
	"rose":         {ID: "rose", Name: "Rose", Icon: assets.URL("/icons/rose.png"), Type: TypeRose, Width: 560, Height: 700},
	"script":       {ID: "script", Name: "zsh", Icon: assets.URL("/icons/iterm.png"), Type: TypeScript, Width: 640, Height: 400},
}

// DockOrder defines which apps appear in the dock and in what order.
var DockOrder = []string{
	"finder", "terminal", "obsidian", "slack", "spotify", "steam",
	"devin", "zed", "xcode", "figma", "hermes", "affinity", "whatsapp", "imsg", "browseros",
}

// AppPathPrefix marks a desktop icon that opens an app rather than a file. The desktop VM carries
// it as the icon's path, so the frontend keeps calling openFile for everything on the desktop.
const AppPathPrefix = "app:"

func Get(id string) (App, bool) {
	a, ok := Registry[id]
	return a, ok
}
