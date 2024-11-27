package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/csat/mock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCSATUsecaseGetStatistics_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	ctx := context.WithValue(context.Background(), utils.RequestIDKey{}, "test-request-id")
	mockRepo := mock.NewMockRepo(ctrl)
	mockLogger := logger.New(&cfg.Service.Logger)

	csatUsecase := NewCSATUsecase(mockRepo, mockLogger)

	mockStats := []*models.CSATStat{
		{ID: 1, Topic: "UX", QuestionID: 101, Question: "How is the design?", AverageScore: 4.5},
		{ID: 2, Topic: "Performance", QuestionID: 102, Question: "Is the app fast?", AverageScore: 4.0},
	}
	mockRepo.EXPECT().GetStatistics(ctx).Return(mockStats, nil)

	statsDTO, err := csatUsecase.GetStatistics(ctx)

	require.NoError(t, err)
	require.Len(t, statsDTO, 2)
	require.Equal(t, "UX", statsDTO[0].Topic)
	require.Equal(t, "Performance", statsDTO[1].Topic)
}

func TestCSATUsecaseGetStatistics_DBError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	ctx := context.WithValue(context.Background(), utils.RequestIDKey{}, "test-request-id")
	mockRepo := mock.NewMockRepo(ctrl)
	mockLogger := logger.New(&cfg.Service.Logger)

	csatUsecase := NewCSATUsecase(mockRepo, mockLogger)

	mockRepo.EXPECT().GetStatistics(ctx).Return(nil, errors.New("database error"))

	statsDTO, err := csatUsecase.GetStatistics(ctx)

	require.Error(t, err)
	require.Nil(t, statsDTO)
}

func TestCSATUsecaseGetQuestionsByTopic_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	ctx := context.WithValue(context.Background(), utils.RequestIDKey{}, "test-request-id")
	mockRepo := mock.NewMockRepo(ctrl)
	mockLogger := logger.New(&cfg.Service.Logger)

	csatUsecase := NewCSATUsecase(mockRepo, mockLogger)

	topic := "UX"
	mockQuestions := []*models.CSATQuestion{
		{ID: 1, Question: "How is the UI?"},
		{ID: 2, Question: "Is navigation easy?"},
	}
	mockRepo.EXPECT().GetQuestionsByTopic(ctx, topic).Return(mockQuestions, nil)

	questionsDTO, err := csatUsecase.GetQuestionsByTopic(ctx, topic)

	require.NoError(t, err)
	require.Len(t, questionsDTO, 2)
	require.Equal(t, "How is the UI?", questionsDTO[0].Question)
	require.Equal(t, "Is navigation easy?", questionsDTO[1].Question)
}

func TestCSATUsecaseGetQuestionsByTopic_DBError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	ctx := context.WithValue(context.Background(), utils.RequestIDKey{}, "test-request-id")
	mockRepo := mock.NewMockRepo(ctrl)
	mockLogger := logger.New(&cfg.Service.Logger)

	csatUsecase := NewCSATUsecase(mockRepo, mockLogger)

	topic := "UX"
	mockRepo.EXPECT().GetQuestionsByTopic(ctx, topic).Return(nil, errors.New("database error"))

	questionsDTO, err := csatUsecase.GetQuestionsByTopic(ctx, topic)

	require.Error(t, err)
	require.Nil(t, questionsDTO)
}

func TestCSATUsecaseSubmitAnswer_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	ctx := context.WithValue(context.Background(), utils.RequestIDKey{}, "test-request-id")
	mockRepo := mock.NewMockRepo(ctrl)
	mockLogger := logger.New(&cfg.Service.Logger)

	csatUsecase := NewCSATUsecase(mockRepo, mockLogger)

	mockAnswer := &models.CSATAnswer{
		Score:          5,
		UserID:         uuid.New(),
		CSATQuestionID: 101,
	}
	mockRepo.EXPECT().InsertAnswer(ctx, mockAnswer).Return(mockAnswer, nil)

	answerDTO, err := csatUsecase.SubmitAnswer(ctx, mockAnswer)

	require.NoError(t, err)
	require.NotNil(t, answerDTO)
	require.Equal(t, mockAnswer.Score, answerDTO.Score)
}

func TestCSATUsecaseSubmitAnswer_DBError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	ctx := context.WithValue(context.Background(), utils.RequestIDKey{}, "test-request-id")
	mockRepo := mock.NewMockRepo(ctrl)
	mockLogger := logger.New(&cfg.Service.Logger)

	csatUsecase := NewCSATUsecase(mockRepo, mockLogger)

	mockAnswer := &models.CSATAnswer{
		Score:          5,
		UserID:         uuid.New(),
		CSATQuestionID: 101,
	}
	mockRepo.EXPECT().InsertAnswer(ctx, mockAnswer).Return(nil, errors.New("database error"))

	answerDTO, err := csatUsecase.SubmitAnswer(ctx, mockAnswer)

	require.Error(t, err)
	require.Nil(t, answerDTO)
}
