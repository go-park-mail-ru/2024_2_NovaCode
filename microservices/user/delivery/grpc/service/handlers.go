package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/user/dto"
	userService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/user"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *usersService) FindByID(ctx context.Context, request *userService.FindByIDRequest) (*userService.FindByIDResponse, error) {
	userUUID, err := uuid.Parse(request.GetUuid())
	if err != nil {
		service.logger.Errorf("failed to parse uuid: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse uuid: %v", err)
	}

	user, err := service.usecase.GetByID(ctx, userUUID)
	if err != nil {
		service.logger.Errorf("cannot find user by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "cannot find user by id: %v", err)
	}

	return &userService.FindByIDResponse{User: service.userDTOToProto(user)}, nil
}

func (service *usersService) userDTOToProto(user *dto.UserDTO) *userService.User {
	return &userService.User{
		Uuid:     user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
		Image:    user.Image,
	}
}
