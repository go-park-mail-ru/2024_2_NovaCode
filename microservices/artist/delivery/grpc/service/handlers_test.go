package service

import (
	"context"
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist/dto"
	mocks "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist/mock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	artistService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/artist"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestFindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	mockUsecase := mocks.NewMockUsecase(ctrl)
	service := NewArtistsService(mockUsecase, logger)

	t.Run("successful find", func(t *testing.T) {
		artist := &dto.ArtistDTO{
			ID:      1,
			Name:    "Test Artist",
			Bio:     "Test bio",
			Country: "Test country",
			Image:   "test_image.jpg",
		}

		mockUsecase.EXPECT().View(gomock.Any(), uint64(1)).Return(artist, nil)

		resp, err := service.FindByID(context.Background(), &artistService.FindByIDRequest{Id: 1})
		assert.NoError(t, err)
		assert.Equal(t, &artistService.FindByIDResponse{
			Artist: &artistService.Artist{
				Id:      1,
				Name:    "Test Artist",
				Bio:     "Test bio",
				Country: "Test country",
				Image:   "test_image.jpg",
			},
		}, resp)
	})

	t.Run("not found error", func(t *testing.T) {
		mockUsecase.EXPECT().View(gomock.Any(), uint64(1)).Return(nil, errors.New("artist not found"))

		_, err := service.FindByID(context.Background(), &artistService.FindByIDRequest{Id: 1})
		assert.Error(t, err)
		assert.Equal(t, status.Errorf(codes.NotFound, "cannot find artist by id: artist not found"), err)
	})
}
