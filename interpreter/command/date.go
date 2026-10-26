package command

import (
	"strings"
	"time"

	"interpreter/response"
	"interpreter/session"
	"interpreter/world"
)

// Date prints the time the way the real command does, in the machine's own zone.
func Date(args []string, w world.World, s *session.Session) response.Response {
	now := time.Now()
	return textResult(strings.Replace(
		now.Format("Mon Jan _2 15:04:05 MST 2006"), now.Format("MST"), zoneName(now), 1), s)
}

// zoneName names the machine's own zone. A browser only reports the offset ("UTC-4"), where the
// same moment in New York is EDT, and a menu bar in Queens would say so.
func zoneName(t time.Time) string {
	switch t.Format("MST") {
	case "UTC-4":
		return "EDT"
	case "UTC-5":
		return "EST"
	}
	return t.Format("MST")
}
