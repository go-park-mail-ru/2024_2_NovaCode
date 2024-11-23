package csat

import "net/http"

type Handlers interface {
	SubmitAnswer(response http.ResponseWriter, request *http.Request)
	GetQuestionsByTopic(response http.ResponseWriter, request *http.Request)
	GetStatistics(response http.ResponseWriter, request *http.Request)
}
