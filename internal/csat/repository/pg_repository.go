package repository

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/pkg/errors"
)

type CSATRepository struct {
	db *sql.DB
}

func NewCSATPGRepository(db *sql.DB) *CSATRepository {
	return &CSATRepository{db: db}
}

func (r *CSATRepository) GetQuestionsByTopic(ctx context.Context, topic string) ([]*models.CSATQuestion, error) {
	var questions []*models.CSATQuestion
	rows, err := r.db.QueryContext(ctx, getQuestionsByTopic, topic)
	if err != nil {
		return nil, errors.Wrap(err, "GetQuestionsByTopic.Query")
	}
	defer rows.Close()

	for rows.Next() {
		question := &models.CSATQuestion{}
		err := rows.Scan(
			&question.ID,
			&question.Question,
		)
		if err != nil {
			return nil, errors.Wrap(err, "GetByArtistID.Query")
		}
		questions = append(questions, question)
	}

	return questions, nil
}
