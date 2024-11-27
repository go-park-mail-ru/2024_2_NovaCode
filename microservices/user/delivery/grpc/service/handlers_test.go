package service

import (
	"context"
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/user/dto"
	mocks "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/user/mock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	userService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/user"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
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
	service := NewUsersService(&cfg.Service.Auth, mockUsecase, logger)

	t.Run("successful find", func(t *testing.T) {
		userID, _ := uuid.Parse("00000000-0000-0000-0000-000000000001")
		user := &dto.UserDTO{
			ID:       userID,
			Username: "testuser",
			Email:    "testuser@example.com",
			Image:    "test_image.jpg",
		}

		mockUsecase.EXPECT().GetByID(gomock.Any(), userID).Return(user, nil)

		resp, err := service.FindByID(context.Background(), &userService.FindByIDRequest{Uuid: "00000000-0000-0000-0000-000000000001"})
		assert.NoError(t, err)
		assert.Equal(t, &userService.FindByIDResponse{
			User: &userService.User{
				Uuid:     "00000000-0000-0000-0000-000000000001",
				Username: "testuser",
				Email:    "testuser@example.com",
				Image:    "test_image.jpg",
			},
		}, resp)
	})

	t.Run("invalid uuid error", func(t *testing.T) {
		_, err := service.FindByID(context.Background(), &userService.FindByIDRequest{Uuid: "invalid-uuid"})
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
		assert.Contains(t, status.Convert(err).Message(), "failed to parse uuid")
	})

	t.Run("not found error", func(t *testing.T) {
		userID, _ := uuid.Parse("00000000-0000-0000-0000-000000000001")
		mockUsecase.EXPECT().GetByID(gomock.Any(), userID).Return(nil, errors.New("user not found"))

		_, err := service.FindByID(context.Background(), &userService.FindByIDRequest{Uuid: "00000000-0000-0000-0000-000000000001"})
		assert.Error(t, err)
		assert.Equal(t, codes.NotFound, status.Code(err))
		assert.Equal(t, "cannot find user by id: user not found", status.Convert(err).Message())
	})
}
