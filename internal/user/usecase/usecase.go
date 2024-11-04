package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/user"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/user/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/s3"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/httpErrors"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
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

	requestId := ctx.Value(utils.RequestIdKey{}).(uuid.UUID)

	if foundUser, err := usecase.pgRepo.FindByUsername(ctx, user.Username); foundUser != nil {
		usecase.logger.Warn(fmt.Sprintf("username '%s' is already taken", user.Username), zap.String("request_id", requestId.String()))
		return nil, fmt.Errorf("user with that username already exists")
	} else if err == nil {
		usecase.logger.Error(fmt.Sprintf("error checking username availability: %v", err), zap.String("request_id", requestId.String()))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrUsernameAvailabilityFailed, err)
	}

	if foundUser, err := usecase.pgRepo.FindByEmail(ctx, user.Email); foundUser != nil {
		usecase.logger.Warn(fmt.Sprintf("email '%s' is already taken", user.Email), zap.String("request_id", requestId.String()))
		return nil, fmt.Errorf("user with that email already exists")
	} else if err == nil {
		usecase.logger.Error(fmt.Sprintf("error checking email availability: %v", err), zap.String("request_id", requestId.String()))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrEmailAvailabilityFailed, err)
	}

	if err := user.HashPassword(); err != nil {
		usecase.logger.Error(fmt.Sprintf("error hashing user password: %v", err), zap.String("request_id", requestId.String()))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrHashPasswordFailed, err)
	}

	insertedUser, err := usecase.pgRepo.Insert(ctx, user)
	if err != nil {
		usecase.logger.Error(fmt.Sprintf("error inserting user into repository: %v", err), zap.String("request_id", requestId.String()))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrCreateUserFailed, err)
	}
	usecase.logger.Info(fmt.Sprintf("user '%s' successfully registered", insertedUser.Username), zap.String("request_id", requestId.String()))

	token, err := utils.GenerateJWT(&usecase.cfg.Auth.Jwt, insertedUser)
	if err != nil {
		usecase.logger.Error(fmt.Sprintf("error generating jwt token: %v", err), zap.String("request_id", requestId.String()))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrGenerateTokenFailed, err)
	}

	userTokenDTO := dto.NewUserTokenDTO(insertedUser, token)
	return userTokenDTO, nil
}

func (usecase *userUsecase) Login(ctx context.Context, user *models.User) (*dto.UserTokenDTO, error) {
	requestId := ctx.Value(utils.RequestIdKey{}).(uuid.UUID)
	foundUser, err := usecase.pgRepo.FindByUsername(ctx, user.Username)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("user not found: %v", err), zap.String("request_id", requestId.String()))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrInvalidUsernamePassword, err)
	}
	usecase.logger.Info(fmt.Sprintf("user found: %s", foundUser.Username), zap.String("request_id", requestId.String()))

	if err := foundUser.ComparePasswords(user.Password); err != nil {
		usecase.logger.Warn(fmt.Sprintf("password comparison failed for user '%s': %v", user.Username, err), zap.String("request_id", requestId.String()))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrInvalidUsernamePassword, err)
	}

	token, err := utils.GenerateJWT(&usecase.cfg.Auth.Jwt, foundUser)
	if err != nil {
		usecase.logger.Error(fmt.Sprintf("failed to generate jwt token: %v", err), zap.String("request_id", requestId.String()))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrGenerateTokenFailed, err)
	}

	userTokenDTO := dto.NewUserTokenDTO(foundUser, token)
	return userTokenDTO, nil
}

func (usecase *userUsecase) Update(ctx context.Context, user *models.User) (*dto.UserDTO, error) {
	currentUser, err := usecase.pgRepo.FindByID(ctx, user.UserID)
	if err != nil {
		usecase.logger.Warnf("user not found: %v", err)
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
		usecase.logger.Errorf("error updating user: %v", err)
		return nil, fmt.Errorf("failed to update user")
	}
	usecase.logger.Infof("user '%s' successfully updated", updatedUser.UserID)

	userDTO := dto.NewUserDTO(updatedUser)
	return userDTO, nil
}

func (usecase *userUsecase) UploadImage(ctx context.Context, userID uuid.UUID, file s3.Upload) (*dto.UserDTO, error) {
	user, err := usecase.pgRepo.FindByID(ctx, userID)
	if err != nil {
		usecase.logger.Warnf("user not found: %v", err)
		return nil, fmt.Errorf("user not found")
	}

	uploadInfo, err := usecase.s3Repo.Put(ctx, file)
	if err != nil {
		usecase.logger.Warnf("failed to save user image: %v", err)
		return nil, fmt.Errorf("failed to save user image")
	}

	imageURL := usecase.generateImageURL(file.Bucket, uploadInfo.Key)

	updatedUserDTO, err := usecase.Update(ctx, &models.User{
		UserID: user.UserID,
		Image:  imageURL,
	})
	if err != nil {
		usecase.logger.Warnf("failed to update user model: %v", err)
		return nil, fmt.Errorf("failed to update user model")
	}

	return updatedUserDTO, nil
}

func (usecase *userUsecase) generateImageURL(bucket string, key string) string {
	return fmt.Sprintf("/%s/%s", bucket, key)
}

func (usecase *userUsecase) GetByID(ctx context.Context, userID uuid.UUID) (*dto.UserDTO, error) {
	user, err := usecase.pgRepo.FindByID(ctx, userID)
	if err != nil {
		usecase.logger.Warnf("failed to find user by id '%s': %v", userID, err)
		return nil, fmt.Errorf("failed to find user")
	}

	userDTO := dto.NewUserDTO(user)
	return userDTO, nil
}

func (usecase *userUsecase) GetByUsername(ctx context.Context, username string) (*dto.UserDTO, error) {
	user, err := usecase.pgRepo.FindByUsername(ctx, username)
	if err != nil {
		usecase.logger.Warnf("failed to find user by name '%s': %v", username, err)
		return nil, fmt.Errorf("failed to find user")
	}

	userDTO := dto.NewUserDTO(user)
	return userDTO, nil
}
