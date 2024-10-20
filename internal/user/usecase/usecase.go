package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/jwt"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/user"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/user/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/httpErrors"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
)

type userUsecase struct {
	cfg    *config.AuthConfig
	repo   user.Repo
	logger logger.Logger
}

func NewUserUsecase(cfg *config.AuthConfig, repo user.Repo, logger logger.Logger) user.Usecase {
	return &userUsecase{cfg, repo, logger}
}

func (usecase *userUsecase) Register(ctx context.Context, user *models.User) (*dto.UserTokenDTO, error) {
	if foundUser, err := usecase.repo.FindByUsername(ctx, user.Username); foundUser != nil {
		usecase.logger.Warnf("username '%s' is already taken", user.Username)
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrUsernameAlreadyExists, foundUser)
	} else if err == nil {
		usecase.logger.Errorf("error checking username availability: %v", err)
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrUsernameAvailabilityFailed, err)
	}

	if foundUser, err := usecase.repo.FindByEmail(ctx, user.Email); foundUser != nil {
		usecase.logger.Warnf("email '%s' is already taken", user.Email)
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrEmailAlreadyExists, foundUser)
	} else if err == nil {
		usecase.logger.Errorf("error checking email availability: %v", err)
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrEmailAvailabilityFailed, err)
	}

	if err := user.HashPassword(); err != nil {
		usecase.logger.Errorf("error hashing user password: %v", err)
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrHashPasswordFailed, err)
	}

	insertedUser, err := usecase.repo.Insert(ctx, user)
	if err != nil {
		usecase.logger.Errorf("error inserting user into repository: %v", err)
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrCreateUserFailed, err)
	}
	usecase.logger.Infof("user '%s' successfully registered", insertedUser.Username)

	token, err := jwt.Generate(&usecase.cfg.Jwt, insertedUser)
	if err != nil {
		usecase.logger.Errorf("error generating jwt token: %v", err)
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrGenerateTokenFailed, err)
	}

	userTokenDTO := dto.NewUserTokenDTO(insertedUser, token)
	return userTokenDTO, nil
}

func (usecase *userUsecase) Login(ctx context.Context, user *models.User) (*dto.UserTokenDTO, error) {
	foundUser, err := usecase.repo.FindByUsername(ctx, user.Username)
	if err != nil {
		usecase.logger.Warnf("user not found: %v", err)
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrInvalidUsernamePassword, err)
	}
	usecase.logger.Infof("user found: %s", foundUser.Username)

	if err := foundUser.ComparePasswords(user.Password); err != nil {
		usecase.logger.Warnf("password comparison failed for user '%s': %v", user.Username, err)
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrInvalidUsernamePassword, err)
	}

	token, err := jwt.Generate(&usecase.cfg.Jwt, foundUser)
	if err != nil {
		usecase.logger.Errorf("failed to generate jwt token: %v", err)
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrGenerateTokenFailed, err)
	}

	userTokenDTO := dto.NewUserTokenDTO(foundUser, token)
	return userTokenDTO, nil
}
