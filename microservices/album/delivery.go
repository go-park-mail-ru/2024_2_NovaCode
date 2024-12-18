package album

import "net/http"

type Handlers interface {
	SearchAlbum(response http.ResponseWriter, request *http.Request)
	ViewAlbum(response http.ResponseWriter, request *http.Request)
	GetAll(response http.ResponseWriter, request *http.Request)
	GetAllByArtistID(response http.ResponseWriter, request *http.Request)
	AddFavoriteAlbum(response http.ResponseWriter, request *http.Request)
	DeleteFavoriteAlbum(response http.ResponseWriter, request *http.Request)
	IsFavoriteAlbum(response http.ResponseWriter, request *http.Request)
	GetFavoriteAlbums(response http.ResponseWriter, request *http.Request)
	GetFavoriteAlbumsCount(response http.ResponseWriter, request *http.Request)
	GetAlbumLikesCount(response http.ResponseWriter, request *http.Request)
}
