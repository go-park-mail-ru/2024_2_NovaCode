package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/pkg/errors"
)

type CSATRepository struct {
	db *sql.DB
}

func NewCSATPGRepository(db *sql.DB) *CSATRepository {
	return &CSATRepository{db: db}
}

func (r *CSATRepository) GetStatistics(ctx context.Context) ([]*models.CSATStat, error) {
	rows, err := r.db.QueryContext(ctx, getStatistics)
	if err != nil {
		return nil, errors.Wrap(err, "GetStatistics.Query")
	}
	defer rows.Close()

	var stats []*models.CSATStat
	for rows.Next() {
		stat := &models.CSATStat{}
		err := rows.Scan(
			&stat.ID,
			&stat.Topic,
			&stat.QuestionID,
			&stat.Question,
			&stat.AverageScore,
		)
		if err != nil {
			return nil, errors.Wrap(err, "GetStatistics.Query")
		}
		stats = append(stats, stat)
	}

	return stats, nil
}

func (r *CSATRepository) GetQuestionsByTopic(ctx context.Context, topic string) ([]*models.CSATQuestion, error) {
	var questions []*models.CSATQuestion

	stmt, err := r.db.PrepareContext(ctx, getQuestionsByTopic)
	if err != nil {
		return nil, errors.Wrap(err, "GetQuestionsByTopic.Prepare")
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, topic)
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
			return nil, errors.Wrap(err, "GetQuestionsByTopic.Scan")
		}
		questions = append(questions, question)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "GetQuestionsByTopic.Rows")
	}

	return questions, nil
}

func (r *CSATRepository) InsertAnswer(ctx context.Context, answer *models.CSATAnswer) (*models.CSATAnswer, error) {
	var insertedAnswer models.CSATAnswer

	if err := r.db.QueryRowContext(
		ctx,
		insertAnswer,
		answer.Score,
		answer.UserID,
		answer.CSATQuestionID,
	).Scan(
		&insertedAnswer.Score,
		&insertedAnswer.UserID,
		&insertedAnswer.CSATQuestionID,
	); err != nil {
		return nil, fmt.Errorf("failed to insert csat answer: %w", err)
	}

	return &insertedAnswer, nil
}
