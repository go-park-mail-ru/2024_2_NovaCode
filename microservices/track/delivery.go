package track

import "net/http"

type Handlers interface {
	SearchTrack(response http.ResponseWriter, request *http.Request)
	ViewTrack(response http.ResponseWriter, request *http.Request)
	GetAll(response http.ResponseWriter, request *http.Request)
	GetAllByArtistID(response http.ResponseWriter, request *http.Request)
	GetAllByAlbumID(response http.ResponseWriter, request *http.Request)
	AddFavoriteTrack(response http.ResponseWriter, request *http.Request)
	DeleteFavoriteTrack(response http.ResponseWriter, request *http.Request)
	IsFavoriteTrack(response http.ResponseWriter, request *http.Request)
	GetFavoriteTracks(response http.ResponseWriter, request *http.Request)
	GetFavoriteTracksCount(response http.ResponseWriter, request *http.Request)
	GetTracksFromPlaylist(response http.ResponseWriter, request *http.Request)
	GetPopular(response http.ResponseWriter, request *http.Request)
}
