package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	Route{
		"GetAuthURI",
		"GET",
		"/auth/uri",
		GetAuthURIHandler,
	},
	Route{
		"GetMusicData",
		"GET",
		"/music/{musicName}",
		GetMusicDataHandler,
	},
	Route{
		"GetPlaylistData",
		"GET",
		"/playlist",
		GetPlaylistDataHandler,
	},
	Route{
		"GetInstantMusic",
		"GET",
		"/music",
		GetInstantMusicHandler,
	},
	Route{
		"GetLastPlayedMusics",
		"GET",
		"/music/last",
		GetLastPlayedMusicsHandler,
	},

	// Route{
	// 	"CallbackSpotify",
	// 	"GET",
	// 	"/callback",

	// },
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {

		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router

}
