package track

import "net/http"

type Handlers interface {
	SearchTrack(response http.ResponseWriter, request *http.Request)
	ViewTrack(response http.ResponseWriter, request *http.Request)
	GetAll(response http.ResponseWriter, request *http.Request)
}
