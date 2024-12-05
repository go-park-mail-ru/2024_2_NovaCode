package dto

import (
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/google/uuid"
)

type RegisterDTO struct {
	Role     string `json:"role,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (dto *RegisterDTO) Validate() error {
	if err := models.ValidateUsername(dto.Username); err != nil {
		return err
	}

	if err := models.ValidateEmail(dto.Email); err != nil {
		return err
	}

	if err := models.ValidatePassword(dto.Password); err != nil {
		return err
	}

	return nil
}

func NewUserFromRegisterDTO(registerDTO *RegisterDTO) *models.User {
	return &models.User{
		Role:     registerDTO.Role,
		Username: registerDTO.Username,
		Email:    registerDTO.Email,
		Password: registerDTO.Password,
	}
}

type LoginDTO struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (dto *LoginDTO) Validate() error {
	if err := models.ValidateUsername(dto.Username); err != nil {
		return err
	}

	if err := models.ValidatePassword(dto.Password); err != nil {
		return err
	}

	return nil
}

func NewUserFromLoginDTO(loginDTO *LoginDTO) *models.User {
	return &models.User{
		Username: loginDTO.Username,
		Password: loginDTO.Password,
	}
}

type UpdateDTO struct {
	UserID   uuid.UUID `json:"id,omitempty"`
	Username string    `json:"username,omitempty"`
	Email    string    `json:"email,omitempty"`
}

func (dto *UpdateDTO) Validate() error {
	if err := models.ValidateUsername(dto.Username); err != nil {
		return err
	}

	if err := models.ValidateEmail(dto.Email); err != nil {
		return err
	}

	return nil
}

func NewUserFromUpdateDTO(updateDTO *UpdateDTO) *models.User {
	return &models.User{
		UserID:   updateDTO.UserID,
		Username: updateDTO.Username,
		Email:    updateDTO.Email,
	}
}

type UserDTO struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email,omitempty"`
	Username string    `json:"username,omitempty"`
	Image    string    `json:"image,omitempty"`
}

func NewUserDTO(user *models.User) *UserDTO {
	return &UserDTO{
		user.UserID,
		user.Email,
		user.Username,
		user.Image,
	}
}

type PublicUserDTO struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username,omitempty"`
	Image    string    `json:"image,omitempty"`
}

func NewPublicUserDTO(userDTO *UserDTO) *PublicUserDTO {
	return &PublicUserDTO{
		userDTO.ID,
		userDTO.Username,
		userDTO.Image,
	}
}

type UserTokenDTO struct {
	User  *UserDTO `json:"user"`
	Token string   `json:"token"`
}

func NewUserTokenDTO(user *models.User, token string) *UserTokenDTO {
	return &UserTokenDTO{
		NewUserDTO(user),
		token,
	}
}
