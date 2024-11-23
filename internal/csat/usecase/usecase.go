package usecase

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/csat"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/csat/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
)

type csatUsecase struct {
	csatRepo csat.Repo
	logger   logger.Logger
}

func NewCSATUsecase(csatRepo csat.Repo, logger logger.Logger) csat.Usecase {
	return &csatUsecase{csatRepo, logger}
}

func (usecase *csatUsecase) GetQuestionsByTopic(ctx context.Context, topic string) ([]*dto.CSATQuestionDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})

	questions, err := usecase.csatRepo.GetQuestionsByTopic(ctx, topic)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("cannot retrieve csat questions by topic '%s': %v", topic, err), requestID)
		return nil, fmt.Errorf("cannot retrieve csat questions")
	}
	usecase.logger.Info(fmt.Sprintf("retrieved csat questions by topic '%s'", topic), requestID)

	var questionsDTO []*dto.CSATQuestionDTO
	for _, question := range questions {
		questionDTO := dto.NewCSATQuestionDTO(question)
		questionsDTO = append(questionsDTO, questionDTO)
	}

	return questionsDTO, nil
}

func (usecase *csatUsecase) SubmitAnswer(ctx context.Context, csatAnswer *models.CSATAnswer) (*dto.CSATAnswerDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})

	answer, err := usecase.csatRepo.InsertAnswer(ctx, csatAnswer)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("user have already answered on this question: %v", err), requestID)
		return nil, fmt.Errorf("user have already answered on this question")
	}

	answerDTO := dto.NewCSATAnswerDTO(answer)
	return answerDTO, nil
}
