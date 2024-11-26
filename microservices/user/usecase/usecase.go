package usecase

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/user"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/user/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/s3"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/google/uuid"
)

type userUsecase struct {
	cfg    *userUsecaseConfig
	pgRepo user.PostgresRepo
	s3Repo s3.S3Repo
	logger logger.Logger
}

type userUsecaseConfig struct {
	Auth  *config.AuthConfig
	Minio *config.MinioConfig
}

func NewUserUsecase(authCfg *config.AuthConfig, minioCfg *config.MinioConfig, pgRepo user.PostgresRepo, s3Repo s3.S3Repo, logger logger.Logger) user.Usecase {
	return &userUsecase{
		&userUsecaseConfig{
			authCfg,
			minioCfg,
		},
		pgRepo,
		s3Repo,
		logger,
	}
}

func (usecase *userUsecase) Register(ctx context.Context, user *models.User) (*dto.UserTokenDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	if foundUser, err := usecase.pgRepo.FindByUsername(ctx, user.Username); foundUser != nil {
		usecase.logger.Warn(fmt.Sprintf("username '%s' is already taken", user.Username), requestID)
		return nil, fmt.Errorf("user with that username already exists")
	} else if err == nil {
		usecase.logger.Error(fmt.Sprintf("error checking username availability: %v", err), requestID)
		return nil, fmt.Errorf("failed to check username availability: %w", err)
	}

	if foundUser, err := usecase.pgRepo.FindByEmail(ctx, user.Email); foundUser != nil {
		usecase.logger.Warn(fmt.Sprintf("email '%s' is already taken", user.Email), requestID)
		return nil, fmt.Errorf("user with that email already exists")
	} else if err == nil {
		usecase.logger.Error(fmt.Sprintf("error checking email availability: %v", err), requestID)
		return nil, fmt.Errorf("failed to check email availability: %w", err)
	}

	if err := user.HashPassword(); err != nil {
		usecase.logger.Error(fmt.Sprintf("error hashing user password: %v", err), requestID)
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	insertedUser, err := usecase.pgRepo.Insert(ctx, user)
	if err != nil {
		usecase.logger.Error(fmt.Sprintf("error inserting user into repository: %v", err), requestID)
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	usecase.logger.Infof("user '%s' successfully registered", insertedUser.Username)

	token, err := utils.GenerateJWT(&usecase.cfg.Auth.Jwt, insertedUser)
	if err != nil {
		usecase.logger.Error(fmt.Sprintf("error generating jwt token: %v", err), requestID)
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	userTokenDTO := dto.NewUserTokenDTO(insertedUser, token)
	return userTokenDTO, nil
}

func (usecase *userUsecase) Login(ctx context.Context, user *models.User) (*dto.UserTokenDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	foundUser, err := usecase.pgRepo.FindByUsername(ctx, user.Username)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("user not found: %v", err), requestID)
		return nil, fmt.Errorf("invalid username or password")
	}
	usecase.logger.Infof("user found: %s", foundUser.Username)

	if err := foundUser.ComparePasswords(user.Password); err != nil {
		usecase.logger.Warn(fmt.Sprintf("password comparison failed for user '%s': %v", user.Username, err), requestID)
		return nil, fmt.Errorf("invalid username or password")
	}

	token, err := utils.GenerateJWT(&usecase.cfg.Auth.Jwt, foundUser)
	if err != nil {
		usecase.logger.Error(fmt.Sprintf("failed to generate jwt token: %v", err), requestID)
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	userTokenDTO := dto.NewUserTokenDTO(foundUser, token)
	return userTokenDTO, nil
}

func (usecase *userUsecase) Update(ctx context.Context, user *models.User) (*dto.UserDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	currentUser, err := usecase.pgRepo.FindByID(ctx, user.UserID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("user not found: %v", err), requestID)
		return nil, fmt.Errorf("failed to find user")
	}

	if user.Username != "" {
		currentUser.Username = user.Username
	}
	if user.Email != "" {
		currentUser.Email = user.Email
	}
	if user.Password != "" {
		currentUser.Password = user.Password
	}
	if user.Image != "" {
		currentUser.Image = user.Image
	}

	updatedUser, err := usecase.pgRepo.Update(ctx, currentUser)
	if err != nil {
		usecase.logger.Error(fmt.Sprintf("error updating user: %v", err), requestID)
		return nil, fmt.Errorf("failed to update user")
	}
	usecase.logger.Infof("user '%s' successfully updated", updatedUser.UserID)

	userDTO := dto.NewUserDTO(updatedUser)
	return userDTO, nil
}

func (usecase *userUsecase) UploadImage(ctx context.Context, userID uuid.UUID, file s3.Upload) (*dto.UserDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	user, err := usecase.pgRepo.FindByID(ctx, userID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("user not found: %v", err), requestID)
		return nil, fmt.Errorf("user not found")
	}

	uploadInfo, err := usecase.s3Repo.Put(ctx, file)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("failed to save user image: %v", err), requestID)
		return nil, fmt.Errorf("failed to save user image")
	}

	imageURL := uploadInfo.Key

	updatedUserDTO, err := usecase.Update(ctx, &models.User{
		UserID: user.UserID,
		Image:  imageURL,
	})
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("failed to update user model: %v", err), requestID)
		return nil, fmt.Errorf("failed to update user model")
	}

	return updatedUserDTO, nil
}

func (usecase *userUsecase) GetByID(ctx context.Context, userID uuid.UUID) (*dto.UserDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	user, err := usecase.pgRepo.FindByID(ctx, userID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("failed to find user by id '%s': %v", userID, err), requestID)
		return nil, fmt.Errorf("failed to find user")
	}

	userDTO := dto.NewUserDTO(user)
	return userDTO, nil
}

func (usecase *userUsecase) GetByUsername(ctx context.Context, username string) (*dto.UserDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	user, err := usecase.pgRepo.FindByUsername(ctx, username)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("failed to find user by name '%s': %v", username, err), requestID)
		return nil, fmt.Errorf("failed to find user")
	}

	userDTO := dto.NewUserDTO(user)
	return userDTO, nil
}
