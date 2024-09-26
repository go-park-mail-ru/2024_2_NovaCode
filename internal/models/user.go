package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID    uuid.UUID `json:"user_id," db:"user_id"`
	Role      string    `json:"role" db:"role"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email" validate:"email"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user.Password = string(hashedPassword)
	return nil
}

func (user *User) ComparePasswords(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
