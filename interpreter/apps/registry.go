package apps

type AppType string

const (
	TypeFinder     AppType = "finder"
	TypeTerminal   AppType = "terminal"
	TypeViewer     AppType = "viewer"
	TypePlaceholder AppType = "placeholder"
	TypeEditor     AppType = "editor"
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
	"finder":       {ID: "finder", Name: "Finder", Icon: "/icons/finder.png", Type: TypeFinder, Width: 720, Height: 480},
	"terminal":     {ID: "terminal", Name: "iTerm", Icon: "/icons/iterm.png", Type: TypeTerminal, Width: 640, Height: 400},
	"viewer-image": {ID: "viewer-image", Name: "Preview", Icon: "/icons/preview.png", Type: TypeViewer, Width: 600, Height: 450},
	"viewer-text":  {ID: "viewer-text", Name: "Preview", Icon: "/icons/preview.png", Type: TypeViewer, Width: 600, Height: 450},
	"obsidian":     {ID: "obsidian", Name: "Obsidian", Icon: "/icons/obsidian.png", Type: TypePlaceholder, Width: 800, Height: 500},
	"slack":        {ID: "slack", Name: "Slack", Icon: "/icons/slack.png", Type: TypePlaceholder, Width: 800, Height: 500},
	"spotify":      {ID: "spotify", Name: "Spotify", Icon: "/icons/spotify.png", Type: TypePlaceholder, Width: 800, Height: 500},
	"steam":        {ID: "steam", Name: "Steam", Icon: "/icons/steam.png", Type: TypePlaceholder, Width: 800, Height: 500},
	"devin":        {ID: "devin", Name: "Devin", Icon: "/icons/devin.png", Type: TypeEditor, Width: 900, Height: 600},
	"zed":          {ID: "zed", Name: "Zed", Icon: "/icons/zed.png", Type: TypeEditor, Width: 900, Height: 600},
	"xcode":        {ID: "xcode", Name: "Xcode", Icon: "/icons/xcode.png", Type: TypeEditor, Width: 900, Height: 600},
	"figma":        {ID: "figma", Name: "Figma", Icon: "/icons/figma.png", Type: TypeEditor, Width: 900, Height: 600},
	"hermes":       {ID: "hermes", Name: "Hermes", Icon: "/icons/hermes.png", Type: TypeEditor, Width: 900, Height: 600},
	"affinity":     {ID: "affinity", Name: "Affinity Designer", Icon: "/icons/affinity.png", Type: TypePlaceholder, Width: 900, Height: 600},
	"whatsapp":     {ID: "whatsapp", Name: "WhatsApp", Icon: "/icons/whatsapp.png", Type: TypePlaceholder, Width: 800, Height: 500},
	"imsg":         {ID: "imsg", Name: "Messages", Icon: "/icons/imsg.png", Type: TypePlaceholder, Width: 800, Height: 500},
}

// DockOrder defines which apps appear in the dock and in what order.
var DockOrder = []string{
	"finder", "terminal", "obsidian", "slack", "spotify", "steam",
	"devin", "zed", "xcode", "figma", "hermes", "affinity", "whatsapp", "imsg",
}

func Get(id string) (App, bool) {
	a, ok := Registry[id]
	return a, ok
}
