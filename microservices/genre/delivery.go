package genre

import "net/http"

type Handlers interface {
	GetAll(response http.ResponseWriter, request *http.Request)
	GetAllByArtistID(response http.ResponseWriter, request *http.Request)
	GetAllByTrackID(response http.ResponseWriter, request *http.Request)
}
