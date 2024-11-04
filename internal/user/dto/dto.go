package dto

import (
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/google/uuid"
)

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
