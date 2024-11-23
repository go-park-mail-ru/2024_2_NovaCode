package repository

const (
	getQuestionsByTopic = `
	SELECT q.id, q.title
	FROM csat_question q
		JOIN csat c ON q.csat_id = c.id
	WHERE c.topic = $1`
)
