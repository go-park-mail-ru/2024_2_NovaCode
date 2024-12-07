package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCSATRepositoryGetStatistics_Success(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)

	csatRepo := NewCSATPGRepository(db, logger)

	columns := []string{"id", "topic", "question_id", "question", "average_score"}
	mockRows := sqlmock.NewRows(columns).
		AddRow(1, "UX Design", 101, "How would you rate our UI?", 4.5).
		AddRow(2, "Backend Performance", 102, "How would you rate API speed?", 3.8)

	mock.ExpectQuery(getStatistics).WillReturnRows(mockRows)

	ctx := context.Background()
	stats, err := csatRepo.GetStatistics(ctx)

	require.NoError(t, err)
	require.Len(t, stats, 2)
	require.Equal(t, "UX Design", stats[0].Topic)
	require.Equal(t, "Backend Performance", stats[1].Topic)
}

func TestCSATRepositoryGetStatistics_ConnDone(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)

	csatRepo := NewCSATPGRepository(db, logger)

	mock.ExpectQuery(getStatistics).WillReturnError(sql.ErrConnDone)

	ctx := context.Background()
	stats, err := csatRepo.GetStatistics(ctx)

	require.Error(t, err)
	require.Nil(t, stats)
}

func TestCSATRepositoryGetQuestionsByTopic_Success(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)

	csatRepo := NewCSATPGRepository(db, logger)

	topic := "UX Design"
	columns := []string{"id", "question"}
	mockRows := sqlmock.NewRows(columns).
		AddRow(101, "How would you rate our UI?").
		AddRow(102, "How would you rate our navigation?")

	mock.ExpectPrepare(getQuestionsByTopic).
		ExpectQuery().
		WithArgs(topic).
		WillReturnRows(mockRows)

	ctx := context.Background()
	questions, err := csatRepo.GetQuestionsByTopic(ctx, topic)

	require.NoError(t, err)
	require.Len(t, questions, 2)
	require.Equal(t, "How would you rate our UI?", questions[0].Question)
	require.Equal(t, "How would you rate our navigation?", questions[1].Question)
}

func TestCSATRepositoryGetQuestionsByTopic_NoRows(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)

	csatRepo := NewCSATPGRepository(db, logger)

	topic := "Nonexistent Topic"
	mock.ExpectPrepare(getQuestionsByTopic).
		ExpectQuery().
		WithArgs(topic).
		WillReturnError(sql.ErrNoRows)

	ctx := context.Background()
	questions, err := csatRepo.GetQuestionsByTopic(ctx, topic)

	require.Error(t, err)
	require.Nil(t, questions)
}

func TestCSATRepositoryInsertAnswer_Success(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)

	csatRepo := NewCSATPGRepository(db, logger)

	answer := &models.CSATAnswer{
		Score:          5,
		UserID:         uuid.New(),
		CSATQuestionID: 101,
	}

	mockRows := sqlmock.NewRows([]string{"score", "user_id", "csat_question_id"}).
		AddRow(answer.Score, answer.UserID, answer.CSATQuestionID)

	mock.ExpectQuery(insertAnswer).WithArgs(
		answer.Score,
		answer.UserID,
		answer.CSATQuestionID,
	).WillReturnRows(mockRows)

	ctx := context.Background()
	insertedAnswer, err := csatRepo.InsertAnswer(ctx, answer)

	require.NoError(t, err)
	require.NotNil(t, insertedAnswer)
	require.Equal(t, answer.Score, insertedAnswer.Score)
	require.Equal(t, answer.UserID, insertedAnswer.UserID)
	require.Equal(t, answer.CSATQuestionID, insertedAnswer.CSATQuestionID)
}

func TestCSATRepositoryInsertAnswer_ConnDone(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)

	csatRepo := NewCSATPGRepository(db, logger)

	answer := &models.CSATAnswer{
		Score:          5,
		UserID:         uuid.New(),
		CSATQuestionID: 101,
	}

	mock.ExpectQuery(insertAnswer).WithArgs(
		answer.Score,
		answer.UserID,
		answer.CSATQuestionID,
	).WillReturnError(sql.ErrConnDone)

	ctx := context.Background()
	insertedAnswer, err := csatRepo.InsertAnswer(ctx, answer)

	require.Error(t, err)
	require.Nil(t, insertedAnswer)
}
