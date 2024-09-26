package dto

import "github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"

type UserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewUserDTO(user *models.User) *UserDTO {
	return &UserDTO{
		user.Username,
		user.Email,
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
