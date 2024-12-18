package dto

import (
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/google/uuid"
)

//easyjson:json
type RegisterDTO struct {
	Role     string `json:"role,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func NewUserFromRegisterDTO(registerDTO *RegisterDTO) *models.User {
	return &models.User{
		Role:     registerDTO.Role,
		Username: registerDTO.Username,
		Email:    registerDTO.Email,
		Password: registerDTO.Password,
	}
}

//easyjson:json
type LoginDTO struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func NewUserFromLoginDTO(loginDTO *LoginDTO) *models.User {
	return &models.User{
		Username: loginDTO.Username,
		Password: loginDTO.Password,
	}
}

//easyjson:json
type UpdateDTO struct {
	UserID   uuid.UUID `json:"id,omitempty"`
	Username string    `json:"username,omitempty"`
	Email    string    `json:"email,omitempty"`
}

func NewUserFromUpdateDTO(updateDTO *UpdateDTO) *models.User {
	return &models.User{
		UserID:   updateDTO.UserID,
		Username: updateDTO.Username,
		Email:    updateDTO.Email,
	}
}

//easyjson:json
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

//easyjson:json
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

//easyjson:json
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
