package response

import "encoding/json"

type Type string

const (
	LS      Type = "ls"
	Cat     Type = "cat"
	PWD     Type = "pwd"
	Echo    Type = "echo"
	Env     Type = "env"
	History Type = "history"
	Clear   Type = "clear"
	VarSet  Type = "varset"
	Open    Type = "open"
	Err     Type = "error"
)

type Response struct {
	Type    Type        `json:"type"`
	Payload interface{} `json:"payload"`
	CWD     string      `json:"cwd"`
}

type LSPayload struct {
	Entries    []FSEntry `json:"entries"`
	LongFormat bool      `json:"long,omitempty"`
	HumanSize  bool      `json:"human-readable,omitempty"`
	ShowHidden bool      `json:"show-hidden,omitempty"`
	Recursing  bool      `json:"recursive,omitempty"`
	Color      string    `json:"color,omitempty"`
}

type FSEntry struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Size       int64  `json:"size"`
	Modified   string `json:"modified"`
	Hidden     bool   `json:"hidden"`
	Permission string `json:"permission,omitempty"`
	Owner      string `json:"owner,omitempty"`
	Group      string `json:"group,omitempty"`
}

type CatPayload struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type TextPayload struct {
	Text string `json:"text"`
}

type EnvPayload struct {
	Vars map[string]string `json:"vars"`
}

// LsOptions holds all the ls flags and options
type LsOptions struct {
	ShowHidden    bool
	LongFormat    bool
	HumanSize     bool
	Recursing     bool
	SortByTime    bool
	SortBySize    bool
	ColorMode     string // auto, always, never
	SortByReverse bool   // reverse sort order
	FullTime      bool   // use full time format
	Oneline       bool   // one entry per line
	DirectoryOnly bool   // show directory info only
}

type HistoryPayload struct {
	Entries []string `json:"entries"`
}

type ErrorPayload struct {
	Command string `json:"command"`
	Message string `json:"message"`
}

func Error(cmd, msg, cwd string) Response {
	return Response{
		Type:    Err,
		Payload: ErrorPayload{Command: cmd, Message: msg},
		CWD:     cwd,
	}
}

func Marshal(r Response) string {
	b, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return string(b)
}
