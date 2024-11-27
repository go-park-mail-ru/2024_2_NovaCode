package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album/dto"
	mocks "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album/mock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	albumService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/album"
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
	service := NewAlbumsService(mockUsecase, logger)

	t.Run("successful find", func(t *testing.T) {
		releaseDate := time.Date(2024, time.October, 1, 0, 0, 0, 0, time.UTC)
		album := &dto.AlbumDTO{
			ID:          1,
			Name:        "Test Album",
			ReleaseDate: releaseDate,
			Image:       "test_image.jpg",
		}

		mockUsecase.EXPECT().View(gomock.Any(), uint64(1)).Return(album, nil)

		resp, err := service.FindByID(context.Background(), &albumService.FindByIDRequest{Id: 1})
		assert.NoError(t, err)
		assert.Equal(t, &albumService.FindByIDResponse{
			Album: &albumService.Album{
				Id:          1,
				Name:        "Test Album",
				ReleaseDate: timestamppb.New(releaseDate),
				Image:       "test_image.jpg",
			},
		}, resp)
	})

	t.Run("not found error", func(t *testing.T) {
		mockUsecase.EXPECT().View(gomock.Any(), uint64(1)).Return(nil, errors.New("album not found"))

		_, err := service.FindByID(context.Background(), &albumService.FindByIDRequest{Id: 1})
		assert.Error(t, err)
		assert.Equal(t, status.Errorf(codes.NotFound, "cannot find album by id: album not found"), err)
	})
}
