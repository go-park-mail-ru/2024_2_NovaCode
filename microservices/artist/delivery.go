package artist

import "net/http"

type Handlers interface {
	SearchArtist(response http.ResponseWriter, request *http.Request)
	ViewArtist(response http.ResponseWriter, request *http.Request)
	GetAll(response http.ResponseWriter, request *http.Request)
}
