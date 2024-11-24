package repository

const (
	getStatistics = `
  SELECT 
    csat.id AS topic_id,
    csat.topic AS topic_name,
    csat_question.id AS question_id,
    csat_question.title AS question_title,
    AVG(csat_answer.score) AS average_score
  FROM 
    csat
  JOIN 
    csat_question ON csat.id = csat_question.csat_id
  JOIN 
    csat_answer ON csat_question.id = csat_answer.csat_question_id
  GROUP BY 
    csat.id, csat.topic, csat_question.id, csat_question.title
  `

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
