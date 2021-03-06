package spotify

import (
	cycle "github.com/joaoaneto/radiup/cycle/model"
	"github.com/joaoaneto/spotify"
)

type StreamerSuggestionOperator interface {
	GetUpdatedMusicList(cycleID int, auth spotify.Authenticator) (cycle.StreamerSuggestion, error)
}

/*The method that will receive de signal from the cycle have to implement this interface*/
type CycleListener interface {
	Notified(c *cycle.Cycle)
}

type VoluntarySuggestionOperator interface {
	VerifyUserVote(cycleID int, musicID string, user string) (bool, error)
	HasVoluntarySuggestion(cycleID int, musicID string) bool
}
