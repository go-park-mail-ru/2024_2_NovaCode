package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/pkg/errors"
)

type CSATRepository struct {
	db     *sql.DB
	logger logger.Logger
}

func NewCSATPGRepository(db *sql.DB, logger logger.Logger) *CSATRepository {
	return &CSATRepository{db, logger}
}

func (r *CSATRepository) GetStatistics(ctx context.Context) ([]*models.CSATStat, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	rows, err := r.db.QueryContext(ctx, getStatistics)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[csat repo] failed to query context in GetStatistics: %v", err), requestID)
		return nil, errors.Wrap(err, "GetStatistics.Query")
	}
	r.logger.Info("[csat repo] successful GetStatistics query context", requestID)
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
			r.logger.Error(fmt.Sprintf("[csat repo] failed to scan rows in GetStatistics: %v", err), requestID)
			return nil, errors.Wrap(err, "GetStatistics.Scan")
		}
		r.logger.Info("[csat repo] successful GetStatistics scan rows", requestID)
		stats = append(stats, stat)
	}

	return stats, nil
}

func (r *CSATRepository) GetQuestionsByTopic(ctx context.Context, topic string) ([]*models.CSATQuestion, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	var questions []*models.CSATQuestion

	stmt, err := r.db.PrepareContext(ctx, getQuestionsByTopic)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[csat repo] failed to prepare context in GetQuestionsByTopic: %v", err), requestID)
		return nil, errors.Wrap(err, "GetQuestionsByTopic.Prepare")
	}
	r.logger.Info("[csat repo] successful GetQuestionsByTopic prepare context", requestID)
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, topic)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[csat repo] failed to query context in GetQuestionsByTopic: %v", err), requestID)
		return nil, errors.Wrap(err, "GetQuestionsByTopic.Query")
	}
	r.logger.Info("[csat repo] successful GetQuestionsByTopic query context", requestID)
	defer rows.Close()

	for rows.Next() {
		question := &models.CSATQuestion{}
		err := rows.Scan(
			&question.ID,
			&question.Question,
		)
		if err != nil {
			r.logger.Error(fmt.Sprintf("[csat repo] failed to scan rows in GetQuestionsByTopic: %v", err), requestID)
			return nil, errors.Wrap(err, "GetQuestionsByTopic.Scan")
		}
		r.logger.Info("[csat repo] successful GetQuestionsByTopic scan rows", requestID)
		questions = append(questions, question)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error(fmt.Sprintf("[csat repo] failed to get rows in GetQuestionsByTopic: %v", err), requestID)
		return nil, errors.Wrap(err, "GetQuestionsByTopic.Rows")
	}

	return questions, nil
}

func (r *CSATRepository) InsertAnswer(ctx context.Context, answer *models.CSATAnswer) (*models.CSATAnswer, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
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
		r.logger.Error(fmt.Sprintf("[csat repo] failed to scan row InsertAnswer: %v", err), requestID)
		return nil, fmt.Errorf("InserAnswer.Scan: %w", err)
	}
	r.logger.Info("[csat repo] successful InsertAnswer scan row", requestID)

	return &insertedAnswer, nil
}
