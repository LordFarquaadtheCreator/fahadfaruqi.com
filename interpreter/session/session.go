package session

type Session struct {
	CWD     string
	Vars    map[string]string
	History []string
}

func New(env map[string]string) *Session {
	home := env["HOME"]
	if home == "" {
		home = "~"
	}
	return &Session{
		CWD:     home,
		Vars:    make(map[string]string),
		History: make([]string, 0),
	}
}

func (s *Session) SetVar(k, v string) {
	s.Vars[k] = v
}

func (s *Session) AppendHistory(input string) {
	s.History = append(s.History, input)
}

func (s *Session) SetCWD(path string) {
	s.CWD = path
}
