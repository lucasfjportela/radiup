package business

import (
	"github.com/joaoaneto/radiup/cycle"
	"github.com/zmb3/spotify"
)

type StreamerSuggestionOperator interface {
	GetUpdatedMusicList(auth spotify.Authenticator) (cycle.StreamerSuggestion, error)
}

/*The method that will receive de signal from the cycle have to implement this interface*/
type CycleListener interface {
	Notified()
}