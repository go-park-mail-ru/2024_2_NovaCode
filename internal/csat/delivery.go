package csat

import "net/http"

type Handlers interface {
	SubmitAnswer(response http.ResponseWriter, request *http.Request)
	GetQuestions(response http.ResponseWriter, request *http.Request)
	GetStatistics(response http.ResponseWriter, request *http.Request)
}
