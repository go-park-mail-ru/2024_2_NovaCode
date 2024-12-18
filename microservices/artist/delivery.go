package artist

import "net/http"

type Handlers interface {
	SearchArtist(response http.ResponseWriter, request *http.Request)
	ViewArtist(response http.ResponseWriter, request *http.Request)
	GetAll(response http.ResponseWriter, request *http.Request)
	AddFavoriteArtist(response http.ResponseWriter, request *http.Request)
	DeleteFavoriteArtist(response http.ResponseWriter, request *http.Request)
	IsFavoriteArtist(response http.ResponseWriter, request *http.Request)
	GetFavoriteArtists(response http.ResponseWriter, request *http.Request)
	GetFavoriteArtistsCount(response http.ResponseWriter, request *http.Request)
	GetArtistLikesCount(response http.ResponseWriter, request *http.Request)
}
