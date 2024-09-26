package usecase

import (
	"context"
	"fmt"

	"github.com/daronenko/auth/config"
	"github.com/daronenko/auth/internal/jwt"
	"github.com/daronenko/auth/internal/models"
	"github.com/daronenko/auth/internal/user"
	"github.com/daronenko/auth/internal/user/dto"
	"github.com/daronenko/auth/pkg/logger"
)

type userUsecase struct {
	cfg    *config.Config
	repo   user.Repo
	logger logger.Logger
}

func NewUserUsecase(cfg *config.Config, repo user.Repo, logger logger.Logger) user.Usecase {
	return &userUsecase{cfg, repo, logger}
}

func (usecase *userUsecase) Register(ctx context.Context, user *models.User) (*dto.UserTokenDTO, error) {
	if foundUser, err := usecase.repo.FindByUsername(ctx, user.Username); foundUser != nil {
		usecase.logger.Warnf("username '%s' is already taken", user.Username)
		return nil, fmt.Errorf("user with that username already exists")
	} else if err == nil {
		usecase.logger.Errorf("error checking username availability: %v", err)
		return nil, fmt.Errorf("failed to check username availability: %w", err)
	}

	if foundUser, err := usecase.repo.FindByEmail(ctx, user.Email); foundUser != nil {
		usecase.logger.Warnf("email '%s' is already taken", user.Email)
		return nil, fmt.Errorf("user with that email already exists")
	} else if err == nil {
		usecase.logger.Errorf("error checking email availability: %v", err)
		return nil, fmt.Errorf("failed to check email availability: %w", err)
	}

	if err := user.HashPassword(); err != nil {
		usecase.logger.Errorf("error hashing user password: %v", err)
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	insertedUser, err := usecase.repo.Insert(ctx, user)
	if err != nil {
		usecase.logger.Errorf("error inserting user into repository: %v", err)
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	usecase.logger.Infof("user '%s' successfully registered", insertedUser.Username)

	token, err := jwt.Generate(usecase.cfg, insertedUser)
	if err != nil {
		usecase.logger.Errorf("error generating jwt token: %v", err)
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	userTokenDTO := dto.NewUserTokenDTO(insertedUser, token)
	return userTokenDTO, nil
}

func (usecase *userUsecase) Login(ctx context.Context, user *models.User) (*dto.UserTokenDTO, error) {
	foundUser, err := usecase.repo.FindByUsername(ctx, user.Username)
	if err != nil {
		usecase.logger.Warnf("user not found: %v", err)
		return nil, fmt.Errorf("invalid username or password")
	}
	usecase.logger.Infof("user found: %s", foundUser.Username)

	if err := foundUser.ComparePasswords(user.Password); err != nil {
		usecase.logger.Warnf("password comparison failed for user '%s': %v", user.Username, err)
		return nil, fmt.Errorf("invalid username or password")
	}

	token, err := jwt.Generate(usecase.cfg, foundUser)
	if err != nil {
		usecase.logger.Errorf("failed to generate jwt token: %v", err)
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	userTokenDTO := dto.NewUserTokenDTO(foundUser, token)
	return userTokenDTO, nil
}
