package repository

const (
	getQuestionsByTopic = `
	SELECT q.id, q.title
	FROM csat_question q
		JOIN csat c ON q.csat_id = c.id
	WHERE c.topic = $1`

	insertAnswer = `
	INSERT INTO csat_answer (score, user_id, csat_question_id)
	VALUES ($1, $2, $3)
	RETURNING score, user_id, csat_question_id
	`
)
