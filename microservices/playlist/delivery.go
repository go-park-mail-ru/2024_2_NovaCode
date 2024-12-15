package playlist

import "net/http"

type Handlers interface {
	CreatePlaylist(response http.ResponseWriter, request *http.Request)
	GetAllPlaylists(response http.ResponseWriter, request *http.Request)
	GetPlaylist(response http.ResponseWriter, request *http.Request)
	GetUserPlaylists(response http.ResponseWriter, request *http.Request)
	AddToPlaylist(response http.ResponseWriter, request *http.Request)
	RemoveFromPlaylist(response http.ResponseWriter, request *http.Request)
	DeletePlaylist(response http.ResponseWriter, request *http.Request)
	AddFavoritePlaylist(response http.ResponseWriter, request *http.Request)
	DeleteFavoritePlaylist(response http.ResponseWriter, request *http.Request)
	IsFavoritePlaylist(response http.ResponseWriter, request *http.Request)
	GetFavoritePlaylists(response http.ResponseWriter, request *http.Request)
}
