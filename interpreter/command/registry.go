package command

import (
	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
)

type CommandFunc func(args []string, w world.World, s *session.Session) response.Response

var Registry = map[string]CommandFunc{
	"ls":      LS,
	"cd":      CD,
	"cat":     Cat,
	"pwd":     PWD,
	"echo":    Echo,
	"env":     Env,
	"history": History,
	"clear":   Clear,
	"open":    Open,
	"apps":    Apps,
	"help":    Help,
}
